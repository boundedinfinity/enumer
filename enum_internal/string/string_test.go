package enum_internal_test

import (
	"encoding/json"
	"testing"

	enum_internal "github.com/boundedinfinity/enumer/enum_internal/string"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Parse(t *testing.T) {
	actual, err := enum_internal.MyStringEnum.Parse("myString1")
	expected := enum_internal.MyStringEnum.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Parse_Err(t *testing.T) {
	actual, err := enum_internal.MyStringEnum.Parse("xxxx")
	expected := enum_internal.MyString("")

	assert.ErrorIs(t, err, enum_internal.MyStringEnum.Err)
	assert.Equal(t, err.Error(), "invalid invalid enum_internal.MyString value 'xxxx'. Must be one of myString1, myString2, myString3")
	assert.Equal(t, expected, actual)
}

func Test_String(t *testing.T) {
	actual := enum_internal.MyStringEnum.MyString1.String()
	expected := "myString1"

	assert.Equal(t, expected, actual)
}

func Test_Is(t *testing.T) {
	assert.Equal(t, false, enum_internal.MyStringEnum.Is("thing1000"))
	assert.Equal(t, true, enum_internal.MyStringEnum.Is("myString1"))
}

func Test_Json_Marshal(t *testing.T) {
	input := enum_internal.MyStringEnum.MyString1
	bs, err := json.Marshal(input)

	actual := string(bs)
	expected := `"myString1"`

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Json_Unmarshal(t *testing.T) {
	input := []byte(`"myString1"`)
	var actual enum_internal.MyString

	err := json.Unmarshal(input, &actual)
	expected := enum_internal.MyStringEnum.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Marshal(t *testing.T) {
	input := enum_internal.MyStringEnum.MyString1
	bs, err := yaml.Marshal(input)

	actual := string(bs)
	expected := "myString1\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Unmarshal(t *testing.T) {
	input := []byte("myString1\n")
	var actual enum_internal.MyString

	err := yaml.Unmarshal(input, &actual)
	expected := enum_internal.MyStringEnum.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
