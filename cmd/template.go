package main

// var standaloneTmpl2 = `
// var (
// 	All = []{{ .Name }} {
// 	{{ range $v := .Items }}
// 		{{- title $v }},
// 	{{ end }}
// 	}
// )

// func (t {{ .Name }}) String() string {
// 	return string(t)
// }

// func Parse(v {{ .BaseType }}) ({{ .Name }}, error) {
// 	f, ok := slicer.FindFn(All, func(x {{ .Name }}) bool {
// 		return {{ .Name }}(v) == x
// 	})

// 	if !ok {
// 		return f, ErrorV(v)
// 	}

// 	return f, nil
// }

// func Is(s {{ .BaseType }}) bool {
// 	return slicer.ContainsFn(All, func(v {{ .Name }}) bool {
// 		return {{ .BaseType }}(v) == s
// 	})
// }

// var ErrInvalid = errors.New("invalid enumeration type")

// func ErrorV(v {{ .BaseType }}) error {
// 	return fmt.Errorf(
// 		"%w '%v', must be one of %v",
// 		ErrInvalid, v, slicer.Join(All, ","),
// 	)
// }

// func (t {{ .Name }}) MarshalJSON() ([]byte, error) {
// 	return json.Marshal({{ .BaseType }}(t))
// }

// func (t *{{ .Name }}) UnmarshalJSON(data []byte) error {
// 	var v {{ .BaseType }}

// 	if err := json.Unmarshal(data, &v); err != nil {
// 		return err
// 	}

//     e, err := Parse(v)

//     if err != nil {
//         return err
//     }

// 	*t = e

// 	return nil
// }

// func (t {{ .Name }}) MarshalYAML() (interface{}, error) {
// 	return {{ .BaseType }}(t), nil
// }

// func (t *{{ .Name }}) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	var v {{ .BaseType }}

// 	if err := unmarshal(&v); err != nil {
// 		return err
// 	}

// 	e, err := Parse(v)

// 	if err != nil {
// 		return err
// 	}

// 	*t = e

// 	return nil
// }
// `
