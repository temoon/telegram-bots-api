package main

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

const TelegramBotsApiUrl = "https://core.telegram.org/bots/api"

const BlockNone = ""
const BlockMethods = "methods"
const BlockTypes = "types"

type FindOpts struct {
	level   int
	results int

	Criteria   func(node *html.Node) bool
	MaxLevel   int
	MaxResults int
}

func (opts *FindOpts) IsMaxLevel() bool {
	return opts.MaxLevel != 0 && opts.level > opts.MaxLevel
}

func (opts *FindOpts) IncLevel() {
	opts.level += 1
}

func (opts *FindOpts) DecLevel() {
	opts.level -= 1
}

func (opts *FindOpts) IsMaxResults() bool {
	return opts.MaxResults != 0 && opts.results >= opts.MaxResults
}

func (opts *FindOpts) IncResultsCount() {
	opts.results += 1
}

type Fields map[string]*Field

type Field struct {
	Type        string
	Name        string
	Description string
	Required    bool
}

type Methods map[string]*Method

type Method struct {
	Name        string
	Description string
	Anchor      string
	Returns     string
	Fields      Fields
}

type Types map[string]*Type

type Type struct {
	Name        string
	Description string
	Anchor      string
	Subtypes    []string
	Fields      Fields
}

func main() {
	var err error

	var doc *html.Node
	if doc, err = fetch(TelegramBotsApiUrl); err != nil {
		log.Fatalln(err)
	}

	var methods Methods
	var types Types
	var version string
	if methods, types, version, err = parse(doc); err != nil {
		log.Fatalln(err)
	}

	log.Println(methods, types, version)
}

func fetch(url string) (doc *html.Node, err error) {
	log.Print("Fetching content... ")

	var res *http.Response
	if res, err = http.Get(url); err != nil {
		return
	}
	defer res.Body.Close()

	if doc, err = html.Parse(res.Body); err != nil {
		return
	}

	log.Println("Done!")

	return
}

func parse(doc *html.Node) (methods Methods, types Types, version string, err error) {
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

	// Version
	findVersionOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.ElementNode && node.Data == "p"
		},

		MaxLevel: 1,
	}

	versionNode := findNextNode(doc, &findVersionOpts)
	if versionNode == nil {
		err = errors.New("version not found")
		return
	}

	versionText := getNodeText(versionNode)
	if len(versionText) > 8 { // len("Bot API ") == 8
		version = versionText[8:]
	}

	// Methods and types
	methods = make(Methods)
	types = make(Types)

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
			findA := FindOpts{
				Criteria: func(node *html.Node) bool {
					if node.Type == html.ElementNode && node.Data == "a" {
						attrs := getNodeAttributes(node)
						return attrs["class"] == "anchor" && !strings.Contains(attrs["name"], "-")
					}

					return false
				},
			}

			a := findNextNode(node, &findA)
			if a == nil {
				currentBlock = BlockNone
				currentName = ""

				continue
			}

			attrs := getNodeAttributes(a)
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
					m.Fields = getBlockFields(node)
				}
			case BlockTypes:
				if t, ok := types[currentName]; ok {
					t.Fields = getBlockFields(node)
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

func getBlockFields(node *html.Node) (fields Fields) {
	fields = make(Fields)

	// TODO: ...

	return
}

func getBlockSubtypes(node *html.Node) (types []string) {
	types = make([]string, 0)

	// TODO: ...

	return
}

func getMethodReturnType(desc string) (returns string) {
	re := regexp.MustCompile(`(?:Returns|On success,).*?((?:[Aa]rray of )?[A-Z]\w+)(?: that were sent| of the sent messages?)? (?:object|is returned|on success)`)
	match := re.FindAllStringSubmatch(desc, -1)

	switch match[0][1] {
	case "array of MessageId":
		return "Array of Int"
	case "array of Messages":
		return "Array of Message"
	}

	return
}

func getNodeText(node *html.Node) (text string) {
	findOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.TextNode
		},
	}

	if nodes := findAllNodes(node, &findOpts); nodes != nil {
		for _, n := range nodes {
			text += strings.Trim(n.Data, " ") + " "
		}
	}

	return strings.TrimRight(text, " ")
}

func findNextNode(node *html.Node, opts *FindOpts) *html.Node {
	opts.MaxResults = 1

	nodes := findAllNodes(node, opts)
	if nodes == nil || len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

func findAllNodes(node *html.Node, opts *FindOpts) (nodes []*html.Node) {
	if opts.IsMaxLevel() || opts.IsMaxResults() {
		return
	}

	opts.IncLevel()
	defer opts.DecLevel()

	nodes = make([]*html.Node, 0)

	if opts.Criteria(node) {
		nodes = append(nodes, node)
		opts.IncResultsCount()
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		foundNodes := findAllNodes(child, opts)
		if foundNodes == nil {
			return
		}

		nodes = append(nodes, foundNodes...)
	}

	return
}

func getNodeAttributes(node *html.Node) (attrs map[string]string) {
	attrs = make(map[string]string)

	if node != nil {
		for _, attr := range node.Attr {
			attrs[attr.Key] = attr.Val
		}
	}

	return
}
