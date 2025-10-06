package enum

type Lang int

const (
	JavaScript Lang = iota
	GoLang
	Java
)

func (l Lang) String() string {
	return [...]string{"JavaScript", "GoLang", "Java"}[l]
}

func (l Lang) MainCodeExample() string {
	switch l {
	case JavaScript:
		return `console.log("Hello!");`
	case GoLang:
		return `package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}`
	case Java:
		return `public class Main {
	public static void main(String[] args) {
		System.out.println("Hello!");
	}
}`
	default:
		return ""
	}
}
