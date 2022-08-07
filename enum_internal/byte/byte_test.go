package enum_internal_test

import (
	"encoding/json"
	"testing"

	enum_internal "github.com/boundedinfinity/enumer/enum_internal/byte"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Byte(t *testing.T) {
	actual, err := enum_internal.Parse(1)
	expected := enum_internal.B1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_String(t *testing.T) {
	actual := enum_internal.B1.String()
	expected := "\x01"

	assert.Equal(t, expected, actual)
}

func Test_Is(t *testing.T) {
	assert.Equal(t, false, enum_internal.Is(6))
	assert.Equal(t, true, enum_internal.Is(1))
}

func Test_ErrorV(t *testing.T) {
	actual := enum_internal.ErrorV(5)
	expected := "invalid enumeration type '5', must be one of \x00,\x01,\x02"

	assert.Equal(t, expected, actual.Error())
}

func Test_Json_Marshal(t *testing.T) {
	input := enum_internal.B1
	bs, err := json.Marshal(input)

	actual := string(bs)
	expected := "1"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Json_Unmarshal(t *testing.T) {
	input := []byte("1")
	var actual enum_internal.MyByte

	err := json.Unmarshal(input, &actual)
	expected := enum_internal.B1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Marshal(t *testing.T) {
	input := enum_internal.B1
	bs, err := yaml.Marshal(input)

	actual := string(bs)
	expected := "1\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Yaml_Unmarshal(t *testing.T) {
	input := []byte("1\n")
	var actual enum_internal.MyByte

	err := yaml.Unmarshal(input, &actual)
	expected := enum_internal.B1

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
