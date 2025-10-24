package Languages

import (
	model "pc/Model"
)

func GoLang() []string {
	lang := model.Lang{
		MainCodeExample: `package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}`,
       MainFileName: "main.go",
}
	return []string{lang.MainFileName, lang.MainCodeExample}
}
