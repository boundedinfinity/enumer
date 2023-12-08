package enum_internal_test

import (
	"encoding/json"
	"testing"

	enum_internal "github.com/boundedinfinity/enumer/enum_internal/string"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Parse(t *testing.T) {
	actual, err := enum_internal.MyStrings.Parse("my-string-1")
	expected := enum_internal.MyStrings.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Parse_Err(t *testing.T) {
	actual, err := enum_internal.MyStrings.Parse("xxxx")
	expected := enum_internal.MyString("")

	assert.ErrorIs(t, err, enum_internal.MyStrings.Err)
	// assert.Equal(t, "invalid invalid MyString value 'xxxx'. Must be one of my-string-1, my-string-2, MyString3", err.Error())
	assert.Equal(t, "invalid invalid MyString value 'xxxx'. Must be one of my-string-1, my-string-2, my-string-3", err.Error())
	assert.Equal(t, expected, actual)
}

func Test_String(t *testing.T) {
	assert.Equal(t, "my-string-1", enum_internal.MyStrings.MyString1.String())
	assert.Equal(t, "my-string-2", enum_internal.MyStrings.MyString2.String())
	// assert.Equal(t, "MyString3", enum_internal.MyStrings.MyString3.String())
	assert.Equal(t, "my-string-3", enum_internal.MyStrings.MyString3.String())
}

func Test_Is(t *testing.T) {
	assert.Equal(t, false, enum_internal.MyStrings.Is("thing1000"))
	assert.Equal(t, true, enum_internal.MyStrings.Is("my-string-1"))
}

func Test_Json_Marshal(t *testing.T) {
	input := enum_internal.MyStrings.MyString1
	bs, err := json.Marshal(input)

	actual := string(bs)
	expected := `"my-string-1"`

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Json_Unmarshal(t *testing.T) {
	input := []byte(`"my-string-1"`)
	var actual enum_internal.MyString

	err := json.Unmarshal(input, &actual)
	expected := enum_internal.MyStrings.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Marshal(t *testing.T) {
	input := enum_internal.MyStrings.MyString1
	bs, err := yaml.Marshal(input)

	actual := string(bs)
	expected := "my-string-1\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Unmarshal(t *testing.T) {
	input := []byte("my-string-1\n")
	var actual enum_internal.MyString

	err := yaml.Unmarshal(input, &actual)
	expected := enum_internal.MyStrings.MyString1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
