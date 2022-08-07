package enum_internal

//go:generate ../../enumer -path=./byte.go

type MyByte byte

const (
	B0 MyByte = 0
	B1 MyByte = 1
	B2 MyByte = 2
)
