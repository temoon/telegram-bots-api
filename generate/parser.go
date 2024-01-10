package main

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const TelegramBotsApiUrl = "https://core.telegram.org/bots/api"

const BlockNone = ""
const BlockMethods = "methods"
const BlockTypes = "types"

const InputFileType = "InputFile"
const IoStreamType = "input"

type Types map[string]*Type

func (t Types) GetFilteredNames() (names []string) {
	names = make([]string, 0, len(t))
	for name, value := range t {
		if name == InputFileType || len(value.Subtypes) != 0 {
			continue
		}

		names = append(names, name)
	}

	sort.Strings(names)

	return
}

type Type struct {
	Name        string
	Description string
	Anchor      string
	Subtypes    []string
	Fields      Fields
}

type Methods map[string]*Method

type Method struct {
	Name        string
	Description string
	Anchor      string
	Returns     string
	Fields      Fields
}

type Fields map[string]*Field

func (f Fields) IsMultipart() bool {
	return false // TODO: ...
}

type Field struct {
	Name        string
	Description string
	Type        string
	IsRequired  bool
}

func (f *Field) IsArray() bool {
	return strings.HasPrefix(f.Type, "[]")
}

func (f *Field) IsObject() bool {
	return unicode.IsUpper(rune(f.Type[0]))
}

func (f *Field) IsInputFile() bool {
	return f.Type == IoStreamType
}

func fetch() (doc *html.Node, err error) {
	var res *http.Response
	if res, err = http.Get(TelegramBotsApiUrl); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	if doc, err = html.Parse(res.Body); err != nil {
		return
	}

	return
}

func parse(doc *html.Node) (methods Methods, types Types, err error) {
	// Content
	findContentOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			if node.Type == html.ElementNode && node.Data == "div" {
				attrs := getNodeAttributes(node)
				return attrs["id"] == "dev_page_content"
			}

			return false
		},
	}

	if doc = findNextNode(doc, &findContentOpts); doc == nil {
		err = errors.New("content not found")
		return
	}

	// Methods and types
	methods = make(Methods)
	types = make(Types)

	findAnchor := FindOpts{
		Criteria: func(node *html.Node) bool {
			if node.Type == html.ElementNode && node.Data == "a" {
				attrs := getNodeAttributes(node)
				return attrs["class"] == "anchor" && !strings.Contains(attrs["name"], "-")
			}

			return false
		},
	}

	var currentBlock, currentName string
	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		if node.Type != html.ElementNode {
			continue
		}

		if node.Data == "h3" || node.Data == "hr" {
			currentBlock = BlockNone
			currentName = ""
		}

		if node.Data == "h4" {
			findAnchor.ResetCounters()
			anchor := findNextNode(node, &findAnchor)
			if anchor == nil {
				currentBlock = BlockNone
				currentName = ""

				continue
			}

			attrs := getNodeAttributes(anchor)
			currentName = getNodeText(node)

			if unicode.IsUpper(rune(currentName[0])) {
				currentBlock = BlockTypes

				types[currentName] = &Type{
					Name:     currentName,
					Anchor:   attrs["href"],
					Subtypes: make([]string, 0),
					Fields:   make(Fields),
				}
			} else {
				currentBlock = BlockMethods

				methods[currentName] = &Method{
					Name:   currentName,
					Anchor: attrs["href"],
					Fields: make(Fields),
				}
			}
		}

		if currentBlock == BlockNone || currentName == "" {
			continue
		}

		if node.Data == "p" {
			switch currentBlock {
			case BlockMethods:
				if m, ok := methods[currentName]; ok {
					m.Description += getBlockDescription(node)
				}
			case BlockTypes:
				if t, ok := types[currentName]; ok {
					t.Description += getBlockDescription(node)
				}
			}
		}

		if node.Data == "table" {
			switch currentBlock {
			case BlockMethods:
				if m, ok := methods[currentName]; ok {
					m.Fields = getBlockFields(node, currentBlock)
				}
			case BlockTypes:
				if t, ok := types[currentName]; ok {
					t.Fields = getBlockFields(node, currentBlock)
				}
			}
		}

		if currentBlock == BlockTypes && node.Data == "ul" {
			if t, ok := types[currentName]; ok {
				t.Subtypes = getBlockSubtypes(node)
			}
		}

		if m, ok := methods[currentName]; currentBlock == BlockMethods && ok && m.Description != "" {
			m.Returns = getMethodReturnType(m.Description)
		}
	}

	return
}

func getBlockDescription(node *html.Node) string {
	return getNodeText(node)
}

func getBlockFields(node *html.Node, currentBlock string) (fields Fields) {
	fields = make(Fields)

	findBodyOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.ElementNode && node.Data == "tbody"
		},
	}

	tableBody := findNextNode(node, &findBodyOpts)
	if tableBody == nil {
		return
	}

	findRowsOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.ElementNode && node.Data == "tr"
		},
	}

	tableRows := findAllNodes(tableBody, &findRowsOpts)
	if len(tableRows) == 0 {
		return
	}

	findColsOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.ElementNode && node.Data == "td"
		},
	}

	for _, row := range tableRows {
		findColsOpts.ResetCounters()
		tableCols := findAllNodes(row, &findColsOpts)
		if currentBlock == BlockMethods && len(tableCols) == 4 {
			name := getNodeText(tableCols[0])
			desc := getNodeText(tableCols[3])

			fields[name] = &Field{
				Name:        name,
				Description: desc,
				Type:        correctType(getNodeText(tableCols[1])),
				IsRequired:  getNodeText(tableCols[2]) == "Yes",
			}
		} else if currentBlock == BlockTypes && len(tableCols) == 3 {
			name := getNodeText(tableCols[0])
			desc := getNodeText(tableCols[2])

			fields[name] = &Field{
				Name:        name,
				Description: desc,
				Type:        correctType(getNodeText(tableCols[1])),
				IsRequired:  !strings.HasPrefix(desc, "Optional"),
			}
		} else {
			log.Fatalln("Unexpected number of columns at fields table")
		}
	}

	return
}

func getBlockSubtypes(node *html.Node) (types []string) {
	types = make([]string, 0)

	findItemsOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.ElementNode && node.Data == "li"
		},
	}

	items := findAllNodes(node, &findItemsOpts)
	for _, item := range items {
		types = append(types, getNodeText(item))
	}

	return
}

func getMethodReturnType(desc string) (returns string) {
	re := regexp.MustCompile(`(?:Returns|On success,).*?((?:[Aa]rray of )?[A-Z]\w+)(?: that were sent| of the sent messages?)? (?:object|is returned|on success)`)
	match := re.FindAllStringSubmatch(desc, -1)

	return correctType(match[0][1])
}

func correctType(t string) string {
	tt := strings.Split(t, " or ")

	var item string
	for i := 0; i < len(tt); i++ {
		item = strings.ToLower(tt[i])

		if strings.HasPrefix(item, "array of ") {
			tt[i] = "[]" + correctType(tt[i][9:])
			continue
		}

		switch item {
		case strings.ToLower(InputFileType):
			tt[i] = IoStreamType
		case "messages":
			tt[i] = "Message"
		case "boolean", "true":
			tt[i] = "bool"
		case "float", "float number":
			tt[i] = "float64"
		case "integer":
			tt[i] = "int32"
		case "string":
			tt[i] = "string"
		}
	}

	return strings.Join(tt, " or ")
}
