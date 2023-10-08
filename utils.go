package enumer

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/stringer"
)

func GetName[E ~string]() string {
	return reflect.TypeOf(E("")).String()
}

func IsEq[A ~string, B ~string](a A) func(B) bool {
	return func(b B) bool {
		return string(a) == string(b) ||
			stringer.ToLower(a) == stringer.ToLower(b) ||
			stringer.ToUpper(a) == stringer.ToUpper(b)
	}
}

func MarshalJSON[E ~string](e E) ([]byte, error) {
	return json.Marshal(string(e))
}

func UnmarshalJSON[E ~string](data []byte, e *E, parser func(string) (E, error)) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	p, err := parser(s)

	if err != nil {
		return err
	}

	*e = p

	return nil
}

func MarshalYAML[E ~string](e E) (interface{}, error) {
	return string(e), nil
}

func UnmarshalYAML[E ~string](unmarshal func(interface{}) error, e *E, parser func(string) (E, error)) error {
	var s string

	if err := unmarshal(&s); err != nil {
		return err
	}

	p, err := parser(s)

	if err != nil {
		return err
	}

	*e = p

	return nil
}

func Value[E ~string](e E) (driver.Value, error) {
	return string(e), nil
}

func Scan[E ~string](value interface{}, e *E, parser func(string) (E, error)) error {
	if value == nil {
		return fmt.Errorf("cannot be null")
	}

	dv, err := driver.String.ConvertValue(value)

	if err != nil {
		return err
	}

	s, ok := dv.(string)

	if !ok {
		return fmt.Errorf("not a string")
	}

	p, err := parser(s)

	if err != nil {
		return err
	}

	*e = p

	return nil
}
