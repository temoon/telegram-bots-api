package main

import (
	"context"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

const RefInputFile = "#/components/schemas/InputFile"

const TemplatesDir = "generate/templates/"

const TypesHeaderTemplate = "types_header.tmpl"
const TypesTemplate = "types.tmpl"
const TypesFile = "types.go"

const RequestFileTemplate = "request.tmpl"

type Type struct {
	Name   string
	Fields []TypeField
}

type TypeField struct {
	Key        string
	Name       string
	Type       string
	IsRequired bool
}

type Request struct {
	Imports      []string
	Method       string
	Type         string
	Fields       map[string]RequestField
	HasRequest   bool
	ResponseType string
	IsMultipart  bool
}

type RequestField struct {
	IsRequired    bool
	IsRef         bool
	IsInputFile   bool
	Name          string
	Type          string
	RawType       string
	RawTypeFormat string
	Variants      []RequestField
}

func main() {
	loader := openapi3.NewLoader()

	doc, err := loader.LoadFromFile("generate/swagger.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	if err = doc.Validate(context.Background()); err != nil {
		log.Fatalln(err)
	}

	if err = GenerateTypes(doc); err != nil {
		log.Fatalln(err)
	}

	if err = GenerateRequests(doc); err != nil {
		log.Fatalln(err)
	}
}

func GenerateTypes(doc *openapi3.T) (err error) {
	var file *os.File
	if file, err = os.Create(TypesFile); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	var t *template.Template
	if t, err = template.ParseFiles(TemplatesDir+TypesHeaderTemplate, TemplatesDir+TypesTemplate); err != nil {
		return
	}

	if err = t.ExecuteTemplate(file, TypesHeaderTemplate, nil); err != nil {
		return
	}

	schemas := make([]string, 0, len(doc.Components.Schemas))
	for name := range doc.Components.Schemas {
		schemas = append(schemas, name)
	}
	sort.Strings(schemas)

	for _, name := range schemas {
		schema := doc.Components.Schemas[name]

		required := make(map[string]bool)
		for _, key := range schema.Value.Required {
			required[key] = true
		}

		fields := make([]TypeField, 0, len(schema.Value.Properties))
		for key, value := range schema.Value.Properties {
			field := TypeField{
				Key:        key,
				Name:       strcase.ToCamel(key),
				Type:       GenerateValueType(value, required[key], ""),
				IsRequired: required[key],
			}

			fields = append(fields, field)
		}
		sort.Slice(fields, func(i, j int) bool {
			iRequired := "1"
			if fields[i].IsRequired {
				iRequired = "0"
			}

			jRequired := "1"
			if fields[j].IsRequired {
				jRequired = "0"
			}

			return iRequired+fields[i].Key < jRequired+fields[j].Key
		})

		data := Type{
			Name:   name,
			Fields: fields,
		}

		if err = t.ExecuteTemplate(file, TypesTemplate, data); err != nil {
			return
		}
	}

	return
}

func GenerateRequests(doc *openapi3.T) (err error) {
	var t *template.Template
	if t, err = template.ParseFiles(TemplatesDir + RequestFileTemplate); err != nil {
		return
	}

	for method, path := range doc.Paths {
		response := path.Post.Responses["200"].Value.Content.Get("application/json").Schema.Value.Properties["result"]

		if err = GenerateRequestFile(t, method, path.Post.RequestBody, response); err != nil {
			return
		}
	}

	return
}

func GenerateRequestFile(t *template.Template, method string, requestBody *openapi3.RequestBodyRef, response *openapi3.SchemaRef) (err error) {
	var file *os.File
	if file, err = os.Create("requests/" + strcase.ToSnake(method[1:]) + ".go"); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	data := Request{
		Method:       method[1:],
		Type:         strings.Title(method[1:]),
		Fields:       make(map[string]RequestField),
		HasRequest:   requestBody != nil,
		ResponseType: GenerateValueType(response, true, "telegram"),
	}

	imports := make(map[string]bool)
	if requestBody != nil {
		form := requestBody.Value.Content.Get("application/x-www-form-urlencoded")
		if form == nil {
			data.IsMultipart = true
			form = requestBody.Value.Content.Get("multipart/form-data")
		}

		required := make(map[string]bool)
		for _, key := range form.Schema.Value.Required {
			required[key] = true
		}

		for key, value := range form.Schema.Value.Properties {
			variants := make([]RequestField, 0)

			if value.Value.Type == "integer" {
				imports["strconv"] = true
			} else if len(value.Value.AnyOf) != 0 {
				for _, ref := range value.Value.AnyOf {
					if ref.Ref != "" && ref.Ref != RefInputFile || ref.Value.Type == "array" {
						imports["encoding/json"] = true
					} else if ref.Ref == RefInputFile {
						imports["io"] = true
					} else if ref.Value.Type == "integer" {
						imports["strconv"] = true
					}

					variant := RequestField{
						IsRequired:    true,
						IsRef:         ref.Ref != "" && ref.Ref != RefInputFile,
						IsInputFile:   ref.Ref == RefInputFile,
						Name:          GenerateValueType(ref, true, ""),
						Type:          GenerateValueType(ref, true, "telegram"),
						RawType:       ref.Value.Type,
						RawTypeFormat: ref.Value.Format,
					}

					variants = append(variants, variant)
				}
			} else if value.Ref != "" && value.Ref != RefInputFile || value.Value.Type == "array" {
				imports["encoding/json"] = true
			}

			data.Fields[key] = RequestField{
				IsRequired:    required[key],
				IsRef:         value.Ref != "" && value.Ref != RefInputFile,
				IsInputFile:   value.Ref == RefInputFile,
				Name:          strcase.ToCamel(key),
				Type:          GenerateValueType(value, required[key], "telegram"),
				RawType:       value.Value.Type,
				RawTypeFormat: value.Value.Format,
				Variants:      variants,
			}
		}
	}

	data.Imports = make([]string, 0)
	for str := range imports {
		data.Imports = append(data.Imports, str)
	}
	sort.Strings(data.Imports)

	if err = t.ExecuteTemplate(file, RequestFileTemplate, &data); err != nil {
		return
	}

	return
}

func GenerateValueType(value *openapi3.SchemaRef, isRequired bool, pkg string) (t string) {
	if value.Ref != "" {
		path := strings.Split(value.Ref, "/")

		t = path[len(path)-1]
		if pkg != "" {
			t = pkg + "." + t
		}

		if !isRequired {
			t = "*" + t
		}
	} else {
		switch value.Value.Type {
		case "boolean":
			t = "bool"
		case "number":
			t = "float64"
		case "integer":
			switch value.Value.Format {
			case "int64":
				t = "int64"
			default:
				t = "int32"
			}
		case "string":
			t = value.Value.Type
		case "array":
			t = "[]" + GenerateValueType(value.Value.Items, true, pkg)
		default:
			t = "interface{}"
		}
	}

	return
}
