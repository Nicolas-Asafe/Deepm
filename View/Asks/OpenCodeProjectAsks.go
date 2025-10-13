package asks

import (
	"fmt"
	"pc/Kernel/project-manager"
)

func OpenCodeProjectAsks() bool {
	var nameProject string
	var langOption int
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
	fmt.Print("> ")
	fmt.Scanln(&langOption)

	fmt.Println("\nChoose how to open:")
	fmt.Println("[1] VS Code")
	fmt.Println("[2] File Explorer (system default)")
	fmt.Print("> ")
	fmt.Scanln(&openOption)

	return project.OpenCodeProjectCommand(nameProject, langOption, openOption)
}