package enum_internal

//go:generate ../../enumer -path=./string.go

type MyString string

const (
	S1 MyString = "s1"
	S2 MyString = "s2"
	S3 MyString = "s3"
)
