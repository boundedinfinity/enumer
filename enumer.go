package enumer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/boundedinfinity/commons/slices"
)

type V interface {
	~string
}

type Enum[T V] []T

func New[T V](ts ...T) Enum[T] {
	var enum Enum[T]

	for _, t := range ts {
		enum = append(enum, t)
	}

	return enum
}

func (ts Enum[T]) ToString() []string {
	var ss []string

	for _, t := range ts {
		ss = append(ss, string(t))
	}

	return ss
}

func (ts Enum[T]) String() string {
	return fmt.Sprintf("%v", ts.ToString())
}

func (ts Enum[T]) Contains(s string) bool {
	return slices.Contains[T](ts, T(s))
}

var ErrNotInEnumeration = errors.New("not in enumeration")

func errNotInEnumerationf[T V](ts Enum[T], s string) error {
	ss := strings.Join(ts.ToString(), ",")
	return fmt.Errorf("'%v' %w, must be one of %v", s, ErrNotInEnumeration, ss)
}

func (ts Enum[T]) Parse(s string) (T, error) {
	var zero T

	for _, t := range ts {
		if string(t) == s {
			return t, nil
		}
	}

	return zero, errNotInEnumerationf(ts, s)
}
