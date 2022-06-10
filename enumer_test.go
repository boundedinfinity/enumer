package enumer_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/boundedinfinity/enumer"
)

func Test_String(t *testing.T) {
	actual := enum.String()
	expected := "[string1 string2 string3]"
	assert.Equal(t, actual, expected)
}

func Test_Contains(t *testing.T) {
	expected := true
	actual := enum.Contains("string1")

	assert.Equal(t, actual, expected)
}

func Test_Contains_invalid(t *testing.T) {
	expected := false
	actual := enum.Contains("stringx")

	assert.Equal(t, actual, expected)
}

func Test_Parse(t *testing.T) {
	expected := String1
	actual, err := enum.Parse("string1")

	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func Test_Parse_err(t *testing.T) {
	_, err := enum.Parse("stringx")

	assert.NotNil(t, err)
	assert.Equal(t, strings.Contains(err.Error(), enumer.ErrNotInEnumeration.Error()), true)
}

type MyString string

var (
	String1 MyString = "string1"
	String2 MyString = "string2"
	String3 MyString = "string3"
	enum             = enumer.New(String1, String2, String3)
)
