package enum_internal

//go:generate ../../enumer -path=./string.enum.go

type MyString string

type myStrings struct {
	MyString1 MyString
	MyString2 MyString
	MyString3 MyString
}
