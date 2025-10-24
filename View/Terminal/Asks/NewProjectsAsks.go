package asks

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/project-manager"
)

func NewProjectsAsks() bool {
	var name string
	var langOption enum.Lang
	var version float64
	var author string

	fmt.Print("Project name: ")
	fmt.Scanln(&name)
	if name == ""{
		fmt.Println("Enter name project")
		return true
	}

	fmt.Println("\nSelect language:")
	fmt.Println("[0] JavaScript")
	fmt.Println("[1] GoLang")
	fmt.Println("[2] Java")
	fmt.Println("[3] Quit")
	fmt.Print("> ")
	fmt.Scanln(&langOption)
    if langOption == 3 {return false}

	fmt.Print("\nProject version (e.g. 1.0): ")
	fmt.Scanln(&version)

	fmt.Print("\nAuthor: ")
	fmt.Scanln(&author)

	fmt.Println("\nCreating project...")
	success := project.NewProjectCommand(name, langOption, version, author)
	if success {
		fmt.Println("Project created successfully!")
	} else {
		fmt.Println("Failed to create project.")
	}
	return true
}
