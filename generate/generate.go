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

const TypeInputFile = "InputFile"
const RefInputFile = "#/components/schemas/InputFile"
const TypesHeaderTemplate = "types_header.tmpl"
const TypesTemplate = "types.tmpl"
const TypesFile = "types.go"
const MethodsHeaderTemplate = "methods_header.tmpl"
const MethodsTemplate = "methods.tmpl"
const MethodsFile = "methods.go"
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

type Method struct {
	Method        string
	Name          string
	HasRequest    bool
	ResponseType  string
	IsResponseRef bool
}

type Request struct {
	Imports     map[string]bool
	Type        string
	Fields      map[string]RequestField
	IsMultipart bool
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

	if err = GenerateMethods(doc); err != nil {
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
	if t, err = template.ParseFiles("generate/"+TypesHeaderTemplate, "generate/"+TypesTemplate); err != nil {
		return
	}

	if err = t.ExecuteTemplate(file, TypesHeaderTemplate, nil); err != nil {
		return
	}

	schemas := make([]string, 0, len(doc.Components.Schemas))
	for name, _ := range doc.Components.Schemas {
		schemas = append(schemas, name)
	}
	sort.Strings(schemas)

	for _, name := range schemas {
		if name == TypeInputFile {
			continue
		}

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
				Type:       GenerateValueType(value, required[key], true),
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

func GenerateMethods(doc *openapi3.T) (err error) {
	var file *os.File
	if file, err = os.Create(MethodsFile); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	var t *template.Template
	if t, err = template.ParseFiles("generate/"+MethodsHeaderTemplate, "generate/"+MethodsTemplate); err != nil {
		return
	}

	imports := []string{
		"github.com/temoon/go-telegram-bots-api/requests",
	}

	if err = t.ExecuteTemplate(file, MethodsHeaderTemplate, imports); err != nil {
		return
	}

	methods := make([]string, 0, len(doc.Paths))
	for method, _ := range doc.Paths {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	for _, method := range methods {
		path := doc.Paths[method]
		response := path.Post.Responses["200"].Value.Content.Get("application/json").Schema.Value.Properties["result"]

		data := Method{
			Method:        method[1:],
			Name:          strings.Title(method[1:]),
			HasRequest:    path.Post.RequestBody != nil,
			ResponseType:  GenerateValueType(response, true, true),
			IsResponseRef: response.Ref != "",
		}

		if err = t.ExecuteTemplate(file, MethodsTemplate, &data); err != nil {
			return
		}
	}

	return
}

func GenerateRequests(doc *openapi3.T) (err error) {
	var t *template.Template
	if t, err = template.ParseFiles("generate/" + RequestFileTemplate); err != nil {
		return
	}

	for method, path := range doc.Paths {
		if err = GenerateRequestFile(t, method, path.Post.RequestBody); err != nil {
			return
		}
	}

	return
}

func GenerateRequestFile(t *template.Template, method string, requestBody *openapi3.RequestBodyRef) (err error) {
	if requestBody == nil {
		return
	}

	var file *os.File
	if file, err = os.Create("requests/" + strcase.ToSnake(method[1:]) + ".go"); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	data := Request{
		Imports: make(map[string]bool),
		Type:    strings.Title(method[1:]),
	}

	form := requestBody.Value.Content.Get("application/x-www-form-urlencoded")
	if form == nil {
		data.IsMultipart = true
		form = requestBody.Value.Content.Get("multipart/form-data")
	}

	required := make(map[string]bool)
	for _, key := range form.Schema.Value.Required {
		required[key] = true
	}

	data.Fields = make(map[string]RequestField)
	for key, value := range form.Schema.Value.Properties {
		variants := make([]RequestField, 0)

		if value.Value.Type == "integer" {
			data.Imports["strconv"] = true
		} else if len(value.Value.AnyOf) != 0 {
			hasRef := false
			for _, ref := range value.Value.AnyOf {
				if ref.Ref != "" && ref.Ref != RefInputFile || ref.Value.Type == "array" {
					if hasRef {
						continue
					}

					hasRef = true
					data.Imports["encoding/json"] = true
				} else if ref.Ref == RefInputFile {
					data.Imports["io"] = true
				} else if ref.Value.Type == "integer" {
					data.Imports["strconv"] = true
				}

				variant := RequestField{
					IsRequired:    true,
					IsRef:         ref.Ref != "" && ref.Ref != RefInputFile,
					IsInputFile:   ref.Ref == RefInputFile,
					Name:          "",
					Type:          "",
					RawType:       ref.Value.Type,
					RawTypeFormat: ref.Value.Format,
				}

				variants = append(variants, variant)
			}
		} else if value.Ref == RefInputFile {
			data.Imports["io"] = true
		} else if value.Ref != "" || value.Value.Type == "array" {
			data.Imports["encoding/json"] = true
		}

		data.Fields[key] = RequestField{
			IsRequired:    required[key],
			IsRef:         value.Ref != "" && value.Ref != RefInputFile,
			IsInputFile:   value.Ref == RefInputFile,
			Name:          strcase.ToCamel(key),
			Type:          GenerateValueType(value, true, false),
			RawType:       value.Value.Type,
			RawTypeFormat: value.Value.Format,
			Variants:      variants,
		}
	}

	if err = t.ExecuteTemplate(file, RequestFileTemplate, &data); err != nil {
		return
	}

	return
}

func GenerateValueType(value *openapi3.SchemaRef, isRequired bool, allowRef bool) (t string) {
	if allowRef && value.Ref != "" {
		path := strings.Split(value.Ref, "/")

		t = path[len(path)-1]
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
			case "int32":
				t = "uint32"
			default:
				t = "uint64"
			}
		case "string":
			t = value.Value.Type
		case "array":
			t = "[]" + GenerateValueType(value.Value.Items, true, allowRef)
		default:
			t = "interface{}"
		}
	}

	return
}
