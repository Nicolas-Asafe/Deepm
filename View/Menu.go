package view

import (
	"fmt"
	"pc/Model"
	asks "pc/View/Asks"
)


func Menu() {
	for {
		fmt.Println("Select one option:")
		fmt.Println()

		options := []model.Option{
			{OptionNumber: 1, OptionString: "New project", Function:asks.NewProjectsAsks},
			{OptionNumber: 2, OptionString: "Delete project", Function: deleteProject},
			{OptionNumber: 3, OptionString: "Code project", Function: codeProject},
			{OptionNumber: 4, OptionString: "Exit", Function: exitProgram},
		}

		shouldContinue, err := RenderOptions(options)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if !shouldContinue {
			fmt.Println("Exiting program...")
			return 
		}
	}
}


func newProject() bool {
	fmt.Println("Creating new project...")
	return true
}

func deleteProject() bool {
	fmt.Println("Deleting project...")
	return true
}

func codeProject() bool {
	fmt.Println("Opening project code editor...")
	return true
}

func exitProgram() bool {
	return false 
}
