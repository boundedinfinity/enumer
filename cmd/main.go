package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/boundedinfinity/asciibox"
	"github.com/boundedinfinity/go-commoner/stringer"
)

const (
	FilePermissions = 0755
)

var (
	Header = []string{
		"===== DO NOT EDIT =====",
		"Any change will be overwritten",
		"Generated by github.com/boundedinfinity/enumer",
	}
)

type argsData struct {
	Path       string
	Header     string
	HeaderFrom string
	Plural     string
	SkipFormat bool
}

type templateData struct {
	Name       string
	Package    string
	Dir        string
	Filename   string
	Path       string
	Items      []string
	BaseType   string
	Header     string
	SkipFormat bool
}

func main() {
	var args argsData

	data := templateData{
		Items: []string{},
	}

	flag.StringVar(&args.Path, "path", "", "The input file used for the enum being generated.")
	flag.StringVar(&args.Plural, "plural", "", "The plural form of the name. Defaults to <name>s.")
	flag.StringVar(&args.Header, "header", "", "The header contents.")
	flag.StringVar(&args.HeaderFrom, "header-from", "", "Path to file containing the header contents.")
	flag.BoolVar(&args.SkipFormat, "skip-format", false, "Skip source formatting.")
	flag.Parse()

	if err := processArgs(args, &data); err != nil {
		handleErr(err)
	}

	if err := getInfo(args, &data); err != nil {
		handleErr(err)
	}

	if err := processHeader(args, &data); err != nil {
		handleErr(err)
	}

	bs, err := processTemplate(data)

	if err != nil {
		handleErr(err)
	}

	if err := processWrite(data, bs); err != nil {
		handleErr(err)
	}
}

func processArgs(args argsData, data *templateData) error {
	if args.Path == "" {
		return errors.New("input cannot be empty")
	}

	if !stringer.EndsWith(args.Path, ".enum.go") {
		return errors.New("must be a .enum.go file")
	}

	data.Path = args.Path
	data.Path = strings.ReplaceAll(data.Path, ".enum.go", ".enum.gen.go")

	data.Dir = filepath.Dir(args.Path)
	data.Filename = filepath.Base(args.Path)
	data.SkipFormat = args.SkipFormat

	return nil
}

func getInfo(args argsData, data *templateData) error {
	src, err := ioutil.ReadFile(args.Path)

	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, data.Path, src, parser.ParseComments)

	if err != nil {
		return err
	}

	ast.Inspect(parsed, func(x ast.Node) bool {
		// fmt.Printf("XXXX: %v = \n", x)

		switch t := x.(type) {
		case *ast.File:
			// fmt.Printf("File: %v\n", t.Name)
			data.Package = t.Name.Name
		case *ast.TypeSpec:
			// fmt.Printf("TypeSpec: %v = %v\n", t.Name, t.Type)
			tname := fmt.Sprintf("%v", t.Type)

			switch tname {
			case "string", "byte", "int":
				if data.BaseType == "" {
					data.BaseType = tname
				} else if data.BaseType != tname {
					msg := fmt.Sprintf("must be of same type was %v, now %v", data.BaseType, tname)
					panic(msg)
				}
			default:
				panic("must be one of string,~int,byte")
			}

			data.Name = t.Name.Name
		case *ast.ValueSpec:
			// fmt.Printf("ValueSpec: %v = %v\n", t.Names, t.Type)

			for _, vname := range t.Names {
				if fmt.Sprintf("%v", t.Type) == data.Name {
					data.Items = append(data.Items, vname.Name)
				}
			}
		}

		return true
	})

	return nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func processWrite(data templateData, bs []byte) error {
	err := os.MkdirAll(data.Dir, FilePermissions)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(data.Path, bs, FilePermissions); err != nil {
		return err
	}

	return nil
}

func processHeader(args argsData, data *templateData) error {
	var headerContent []string

	if args.Header != "" {
		headerContent = []string{args.Header}
	} else if args.HeaderFrom != "" {
		file, err := os.Open(args.HeaderFrom)

		if err != nil {

			return err
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		headerContent = append(headerContent, scanner.Text())

		if err := scanner.Err(); err != nil {
			return err
		}
	} else {
		headerContent = Header
	}

	data.Header = asciibox.Box(
		headerContent,
		asciibox.BoxOptions{
			Alignment: asciibox.Alignment_Left,
		},
	)

	return nil
}

func processTemplate(data templateData) ([]byte, error) {
	funcs := template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}

	t, err := template.New(data.Filename).Funcs(funcs).Parse(standaloneTmpl)

	if err != nil {
		return []byte{}, nil
	}

	var b bytes.Buffer

	if err := t.Execute(&b, data); err != nil {
		return []byte{}, err
	}

	if data.SkipFormat {
		return b.Bytes(), nil
	} else {
		return format.Source(b.Bytes())
	}
}

var standaloneTmpl = `
{{ .Header }}

package {{ .Package }}

import (
	"database/sql/driver"
	"fmt"
	
	"github.com/boundedinfinity/enumer"
	"github.com/boundedinfinity/go-commoner/slicer"
)

{{- $name := .Name }}

// /////////////////////////////////////////////////////////////////
//  {{ .Name }} Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t {{ .Name }}) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  {{ .Name }} JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t {{ .Name }}) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *{{ .Name }}) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, {{ .Name }}Enum.Parse)
}

// /////////////////////////////////////////////////////////////////
//  {{ .Name }} YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t {{ .Name }}) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *{{ .Name }}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, {{ .Name }}Enum.Parse)
}

// /////////////////////////////////////////////////////////////////
//  {{ .Name }} SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t {{ .Name }}) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *{{ .Name }}) Scan(value interface{}) error {
	return enumer.Scan(value, t, {{ .Name }}Enum.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

var {{ .Name }}Enum = {{ lower .Name }}Enum{}

type {{ lower .Name }}Enum struct {
	name  string
	All   []{{ .Name }}
	Err   error
{{- range $v := .Items }}
	{{ title $v }} {{ $name -}}
{{ end }}
}

func init() {
	{{ .Name }}Enum.name = enumer.GetName[{{ .Name }}]()
	{{ .Name }}Enum.Err = fmt.Errorf("invalid %v", {{ .Name }}Enum.name)
	
	{{ range $v := .Items -}}
		{{- $name }}Enum.{{- title $v }} = {{ $name }}("{{- $v }}")
	{{ end }}

	{{ .Name }}Enum.All = []{{ .Name }}{
		{{- range $v := .Items }}
			{{ $name }}Enum.{{ title $v -}},
		{{- end }}
	}
}

func (t {{ lower .Name }}Enum) newErr(a any, values ...{{ .Name }}) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		t.Err, a, slicer.Join(values, ", "),
	)
}

func (t {{ lower .Name }}Enum) ParseFrom(v string, values ...{{ .Name }}) ({{ .Name }}, error) {
	f, ok := slicer.FindFn(values, enumer.IsEq[string, {{ .Name }}](v))

	if !ok {
		return f, t.newErr(v, values...)
	}

	return f, nil
}

func (t {{ lower .Name }}Enum) Parse(v string) ({{ .Name }}, error) {
	return t.ParseFrom(v, t.All...)
}

func (t {{ lower .Name }}Enum) IsFrom(v string, values ...{{ .Name }}) bool {
	return slicer.ContainsFn(values, enumer.IsEq[string, {{ .Name }}](v))
}

func (t {{ lower .Name }}Enum) Is(v string) bool {
	return t.IsFrom(v, t.All...)
}
`
