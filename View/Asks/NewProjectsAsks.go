package asks

import (
	"fmt"
	"pc/Enum"
	"pc/Kernel"
)

func NewProjectsAsks() bool {
	var name string
	var langOption int
	var version float64
	var author string

	fmt.Print("Project name: ")
	fmt.Scanln(&name)

	fmt.Println("\nSelect language:")
	fmt.Println("[0] JavaScript")
	fmt.Println("[1] GoLang")
	fmt.Println("[2] Java")
	fmt.Print("> ")
	fmt.Scanln(&langOption)

	var lang enum.Lang
	switch langOption {
	case 0:
		lang = enum.JavaScript
	case 1:
		lang = enum.GoLang
	case 2:
		lang = enum.Java
	default:
		fmt.Println("Invalid option.")
		return true
	}

	fmt.Print("\nProject version (e.g. 1.0): ")
	fmt.Scanln(&version)

	fmt.Print("\nAuthor: ")
	fmt.Scanln(&author)

	fmt.Println("\nCreating project...")
	success := kernel.NewFileCommand(name, lang, version, author)
	if success {
		fmt.Println("Project created successfully!")
	} else {
		fmt.Println("Failed to create project.")
	}
	return true
}
