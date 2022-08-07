package enum_internal_test

import (
	"encoding/json"
	"testing"

	enum_internal "github.com/boundedinfinity/enumer/enum_internal/string"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Parse(t *testing.T) {
	actual, err := enum_internal.Parse("s1")
	expected := enum_internal.S1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_String(t *testing.T) {
	actual := enum_internal.S1.String()
	expected := "s1"

	assert.Equal(t, expected, actual)
}

func Test_Is(t *testing.T) {
	assert.Equal(t, false, enum_internal.Is("thing1000"))
	assert.Equal(t, true, enum_internal.Is("s2"))
}

func Test_ErrorV(t *testing.T) {
	actual := enum_internal.ErrorV("x")
	expected := "invalid enumeration type 'x', must be one of s1,s2,s3"

	assert.Equal(t, expected, actual.Error())
}

func Test_Json_Marshal(t *testing.T) {
	input := enum_internal.S1
	bs, err := json.Marshal(input)

	actual := string(bs)
	expected := `"s1"`

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Json_Unmarshal(t *testing.T) {
	input := []byte(`"s1"`)
	var actual enum_internal.MyString

	err := json.Unmarshal(input, &actual)
	expected := enum_internal.S1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Marshal(t *testing.T) {
	input := enum_internal.S1
	bs, err := yaml.Marshal(input)

	actual := string(bs)
	expected := "s1\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Unmarshal(t *testing.T) {
	input := []byte("s1\n")
	var actual enum_internal.MyString

	err := yaml.Unmarshal(input, &actual)
	expected := enum_internal.S1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
