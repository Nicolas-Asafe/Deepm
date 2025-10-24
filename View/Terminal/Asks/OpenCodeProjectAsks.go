package asks

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/project-manager"
)

func OpenCodeProjectAsks() bool {
	var nameProject string
	var langOption enum.Lang
	var openOption int

	fmt.Print("Project name: ")
	fmt.Scanln(&nameProject)
	if nameProject == "" {
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

	fmt.Println("\nChoose how to open:")
	fmt.Println("[1] VS Code")
	fmt.Println("[2] File Explorer (system default)")
	fmt.Println("[3] Quit")
	fmt.Print("> ")
	fmt.Scanln(&openOption)
	if langOption == 3 {return false}


	return project.OpenCodeProjectCommand(nameProject, langOption, openOption)
}