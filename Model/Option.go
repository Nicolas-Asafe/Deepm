package model

type Option struct {
	OptionNumber int
	OptionString string
	Function     func() bool
}
