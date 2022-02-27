package enumer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/boundedinfinity/enumer"
)

func TestEnumer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Enumer Suite")
}

type MyString string

var (
	String1 MyString = "string1"
	String2 MyString = "string2"
	String3 MyString = "string3"
	enum             = enumer.New(String1, String2, String3)
)

var _ = Describe("Enumer", func() {
	Context("String()", func() {
		It("should generate string", func() {
			actual := enum.String()
			expected := "[string1 string2 string3]"
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Contains()", func() {
		It("find valid value", func() {
			expected := true
			actual := enum.Contains("string1")

			Expect(actual).To(Equal(expected))
		})

		It("not find invalid value", func() {
			expected := false
			actual := enum.Contains("stringx")

			Expect(actual).To(Equal(expected))
		})
	})

	Context("Parse()", func() {
		It("should parse valid value", func() {
			expected := String1
			actual, err := enum.Parse("string1")

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should not parse invalid value", func() {
			_, err := enum.Parse("stringx")

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring(enumer.ErrNotInEnumeration.Error()))
		})

	})
})
