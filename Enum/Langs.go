package enum

import (
	languages "pc/Languages"
)

type Lang int

const (
	JavaScript Lang = iota
	GoLang
	Java
	None
)

func (l Lang) String() string {
	return [...]string{"JavaScript", "GoLang", "Java", "None"}[l]
}

func GetAllLanguages() []Lang{
	return []Lang{
		JavaScript,GoLang,Java,
	}
}
func GetAllLanguagesStrings() []string{
	return []string{
		JavaScript.String(),GoLang.String(),Java.String(),
	}
}

func (l Lang) MainCodeExample() []string {
	switch l {
	case JavaScript:
		return languages.JavaScriptLang()
	case GoLang:
		return languages.GoLang()
	case Java:
		return languages.JavaLang()
	default:
		return []string{"", ""}
	}
}
