package enum_internal_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/enumer/enum_internal"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Parse(t *testing.T) {
	actual, err := enum_internal.Parse("thing1")
	expected := enum_internal.Thing1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_String(t *testing.T) {
	actual := enum_internal.Thing1.String()
	expected := "thing1"

	assert.Equal(t, expected, actual)
}

func Test_Is(t *testing.T) {
	assert.Equal(t, false, enum_internal.Is("thing1000"))
	assert.Equal(t, true, enum_internal.Is("thing1"))
}

func Test_ErrorV(t *testing.T) {
	actual := enum_internal.ErrorV("x")
	expected := "invalid enumeration type 'x', must be one of thing1,thing2,thing3"

	assert.Equal(t, expected, actual.Error())
}

func Test_Json_Marshal(t *testing.T) {
	input := enum_internal.Thing1
	bs, err := json.Marshal(input)

	actual := string(bs)
	expected := `"thing1"`

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Json_Unmarshal(t *testing.T) {
	input := []byte(`"thing1"`)
	var actual enum_internal.MyEnum

	err := json.Unmarshal(input, &actual)
	expected := enum_internal.Thing1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Marshal(t *testing.T) {
	input := enum_internal.Thing1
	bs, err := yaml.Marshal(input)

	actual := string(bs)
	expected := "thing1\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Unmarshal(t *testing.T) {
	input := []byte("thing1\n")
	var actual enum_internal.MyEnum

	err := yaml.Unmarshal(input, &actual)
	expected := enum_internal.Thing1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
