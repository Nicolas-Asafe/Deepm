package Model

type Option struct {
	OptionNumber int
	OptionString string
	Function     func() bool
}
