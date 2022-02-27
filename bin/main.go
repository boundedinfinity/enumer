package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/boundedinfinity/asciibox"
	"github.com/boundedinfinity/caser"
)

type argsData struct {
	Name       string
	Package    string
	Dir        string
	Filename   string
	Suffix     string
	Prefix     string
	Items      string
	ItemsFrom  string
	Header     string
	HeaderFrom string
	Plural     string
	Standalone bool
}

type templateData struct {
	Name       string
	Package    string
	Dir        string
	Filename   string
	Path       string
	Suffix     string
	Prefix     string
	Items      map[string]string
	Header     string
	Plural     string
	Standalone bool
}

const (
	FilePermissions = 0755
)

var (
	Header = []string{
		"===== DO NOT EDIT =====",
		"Any change will be overwritten",
		"Generated by github.com/boundedinfinity/enumer",
	}

	startsWithDigit = regexp.MustCompile(`^\d`)
)

func processName(args argsData, data *templateData) error {
	if args.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	data.Name = args.Name

	if args.Plural == "" {
		data.Plural = fmt.Sprintf("%vs", args.Name)
	} else {
		data.Plural = args.Plural
	}

	return nil
}

func processItems(args argsData, data *templateData) error {
	if args.Items == "" && args.ItemsFrom == "" {
		return fmt.Errorf("both of items and items-from cannot be empty")
	}

	if args.Items != "" {
		for _, v := range strings.Split(args.Items, ",") {
			v = strings.TrimSpace(v)
			k := normalizeName(v)
			data.Items[k] = v
		}
	} else {
		file, err := os.Open(args.ItemsFrom)

		if err != nil {
			return err
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			v := scanner.Text()
			v = strings.TrimSpace(v)

			if v == "" {
				continue
			}

			k := normalizeName(v)
			data.Items[k] = v
		}

		if err := scanner.Err(); err != nil {
			return err
		}
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

func processPackage(args argsData, data *templateData) error {
	if args.Package == "" {
		data.Package = "main"
	} else {
		data.Package = args.Package
	}

	return nil
}

func processPath(args argsData, data *templateData) error {
	data.Standalone = args.Standalone

	if args.Dir != "" {
		data.Dir = args.Dir
	} else {
		if args.Standalone {
			cwd, err := os.Getwd()

			if err != nil {
				return err
			}

			base := filepath.Base(cwd)

			if base != args.Package {
				data.Dir = filepath.Join(".", args.Package)
			}
		} else {
			data.Dir = "."
		}
	}

	if args.Standalone {
		data.Filename = fmt.Sprintf("%v.enumer.go", args.Name)
	}

	data.Path = filepath.Join(data.Dir, data.Filename)

	return nil
}

func processTemplate(data templateData) ([]byte, error) {
	funcs := template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}

	var t *template.Template
	var err error

	if data.Standalone {
		t, err = template.New(data.Filename).Funcs(funcs).Parse(standaloneTmpl)
	} else {
		t, err = template.New(data.Filename).Funcs(funcs).Parse(embeddedTmpl)
	}

	if err != nil {
		return []byte{}, nil
	}

	var b bytes.Buffer

	if err := t.Execute(&b, data); err != nil {
		return []byte{}, nil
	}

	return format.Source(b.Bytes())
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

func main() {
	var args argsData

	data := templateData{
		Items: make(map[string]string),
	}

	flag.StringVar(&args.Name, "name", "", "The name of the enumeration.")
	flag.StringVar(&args.Package, "package", "", "The package of the enumeration.")
	flag.StringVar(&args.Dir, "dir", "", "The output directory of the enumeration.")
	flag.StringVar(&args.Filename, "filename", "", "The output path used for the optinoal being generated.")
	flag.StringVar(&args.Plural, "plural", "", "The plural form of the name. Defaults to <name>s.")
	flag.BoolVar(&args.Standalone, "standalone", false, "If true, generated in it's on package.")
	flag.StringVar(&args.Items, "items", "", "The comma separated list of items.")
	flag.StringVar(&args.ItemsFrom, "items-from", "", "Path to file containing one item one per line.")
	flag.StringVar(&args.Header, "header", "", "The header contents.")
	flag.StringVar(&args.HeaderFrom, "header-from", "", "Path to file containing the header contents.")
	flag.Parse()

	if err := processName(args, &data); err != nil {
		handleErr(err)
	}

	if err := processItems(args, &data); err != nil {
		handleErr(err)
	}

	if err := processHeader(args, &data); err != nil {
		handleErr(err)
	}

	if err := processPackage(args, &data); err != nil {
		handleErr(err)
	}

	if err := processPath(args, &data); err != nil {
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

func handleErr(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func normalizeName(v string) string {
	o := v

	o = strings.ReplaceAll(o, "-", " ")
	o = strings.ReplaceAll(o, "_", " ")
	o = strings.ReplaceAll(o, "/", " ")
	o = strings.ReplaceAll(o, ".", " ")
	o = strings.ReplaceAll(o, "+", " ")
	o = strings.TrimSpace(o)

	if startsWithDigit.MatchString(o) {
		o = fmt.Sprintf("_%v", o)
	}

	o = caser.PhraseToCamel(o)

	return o
}

var standaloneTmpl = `
{{ .Header }}

package {{ .Package }}

import (
	"fmt"
	"errors"
	"encoding/json"
	"strings"
)

type {{ .Name }} string
type {{ .Plural }} []{{ .Name }}

func Slice(es ...{{ .Name }}) {{ .Plural }} {
	var s {{ .Plural }}

	for _, e := range es {
		s = append(s, e)
	}
	
	return s
}

const (
	{{ range $k, $v := .Items }}
		{{- title $k }} {{ $.Name }} = "{{ $v }}"
	{{ end }}
)

var (
	All = {{ .Plural }} {
	{{ range $k, $v := .Items }}
		{{- title $k }},
	{{ end }}
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) ({{ .Name }}, error) {
	return All.Parse(v)
}

func Strings() ([]string) {
	return All.Strings()
}

func (t {{ .Name }}) String() string {
	return string(t)
}

var Err{{ .Name }}Invalid = errors.New("invalid enumeration type")

func Error(vs {{ .Plural }}, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		Err{{ .Name }}Invalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t {{ .Plural }}) Strings() []string {
	var ss []string
	
	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t {{ .Plural }}) Parse(v string) ({{ .Name }}, error) {
	var o {{ .Name }}
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t {{ .Plural }}) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t {{ .Plural }}) Contains(v {{ .Name }}) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t {{ .Name }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *{{ .Name }}) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

    e, err := Parse(s)

    if err != nil {
        return err
    }

	*t = e

	return nil
}



`

var embeddedTmpl = `
{{ .Header }}

package {{ .Package }}

import (
	"fmt"
	"encoding/json"
	"strings"
)

type {{ .Name }} string

const (
{{ range $k, $v := .Items }}
	{{- $.Name }}_{{ title $k }} {{ $.Prefix }}{{ $.Name }}{{ $.Suffix }} = "{{ $v }}"
{{ end }}
)

var (
	{{ .Name }}List = []{{ .Name }}{
	{{ range $k, $v := .Items }}
		{{- $.Name }}_{{ title $k }},
	{{ end }}
	}
)

func Is{{ .Name }}(v string) bool {
	var f bool

	for _, e := range {{ .Name }}List {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func {{ .Name }}Parse(v string) ({{ .Name }}, error) {
	var o {{ .Name }}
	var f bool
	n := strings.ToLower(v)

	for _, e := range {{ .Name }}List {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Err{{ .Name }}Invalid(v)
	}

	return o, nil
}

func {{ .Name }}ListToStrings(vs []{{ .Name }}) ([]string) {
	var ss []string
	
	for _, v := range vs {
		ss = append(ss, v.String())
	}

	return ss
}

func Err{{ .Name }}Invalid(v string) error {
	var ss []string

	for _, e := range {{ .Name }}List {
		ss = append(ss, string(e))
	}
	
	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(ss, ","),
	)
}

func (t {{ .Name }}) String() string {
	return string(t)
}

func (t {{ .Name }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *{{ .Name }}) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

    e, err := {{ .Name }}Parse(s)

    if err != nil {
        return err
    }

	*t = e

	return nil
}
`