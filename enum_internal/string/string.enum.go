package enum_internal

//go:generate ../../enumer -path=./string.enum.go

type MyString string

const (
	myString1 MyString = ""
	myString2 MyString = ""
	myString3 MyString = ""
)
