package project

import (
	"fmt"
	enum "pc/Enum"
	commands "pc/Kernel/files-commands"
)

func OpenCodeProjectCommand(nameProject string, langOption int, openOption int) bool {
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

	path := fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\%s", langProject,nameProject)
	fmt.Println("\nOpening project...")

	switch openOption {
	case 1:
		if err := commands.OpenVSCode(path); err != nil {
			fmt.Println("Failed to open project in VS Code:", err)
			return true
		}
	case 2:
		if err := commands.OpenInWindows(path); err != nil {
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