package asks

import (
	"fmt"
	"pc/Enum"
	"pc/Kernel/files-manager"
)

func OpenCodeProjectAsks() bool {
	var nameProject string
	var langOption int


	fmt.Print("Project name: ")
	fmt.Scanln(&nameProject)
	if nameProject == ""{
		fmt.Println("Enter name project")
		return true
	}

	fmt.Println("\nSelect language:")
	fmt.Println("[0] JavaScript")
	fmt.Println("[1] GoLang")
	fmt.Println("[2] Java")
	fmt.Print("> ")
	fmt.Scanln(&langOption)
	var langProject enum.Lang
	switch langOption {
	case 0:
		langProject = enum.JavaScript
	case 1:
		langProject = enum.GoLang
	case 2:
		langProject = enum.Java
	default:
		fmt.Println("Invalid language option.")
		return true
	}

	var path string
	switch langProject {
	case enum.JavaScript:
		path = fmt.Sprintf("C:\\Dev\\Projects\\JavaScriptProjects\\%s", nameProject)
	case enum.GoLang:
		path = fmt.Sprintf("C:\\Dev\\Projects\\GoProjects\\%s", nameProject)
	case enum.Java:
		path = fmt.Sprintf("C:\\Dev\\Projects\\JavaProjects\\%s", nameProject)
	default:
		fmt.Println("Unknown language.")
		return true
	}

	fmt.Println("\nChoose how to open:")
	fmt.Println("[1] VS Code")
	fmt.Println("[2] File Explorer (system default)")
	fmt.Print("> ")

	var option int
	fmt.Scanln(&option)

	fmt.Println("\nOpening project...")

	switch option {
	case 1:
		if err := files.OpenVSCode(path); err != nil {
			fmt.Println("Failed to open project in VS Code:", err)
			return true
		}
	case 2:
		if err := files.OpenInWindows(path); err != nil {
			fmt.Println("Failed to open project in File Explorer:", err)
			return true
		}
	default:
		fmt.Println("Invalid option.")
		return true
	}

	fmt.Println("Project opened successfully!")
	return true
}
