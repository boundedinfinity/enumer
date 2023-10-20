package enum_internal

//go:generate ../../enumer -path=./string.enum.go

type MyString string

type myStrings struct {
	MyString1 MyString `json:"my-string-1" enum:"my-string-1"`
	MyString2 MyString `enum:"my-string-2"`
	MyString3 MyString
}
