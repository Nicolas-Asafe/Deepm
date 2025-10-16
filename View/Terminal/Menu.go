package view

import (
	"fmt"
	"pc/Model"
	asks "pc/View/Terminal/Asks"
)


func Menu() {
	fmt.Println("Welcome to Deepm Terminal")
	fmt.Println("")
	
	fmt.Println("Press any key for continue")
	fmt.Scanln()
	for {
		fmt.Println("Select one option:")
		fmt.Println()

		options := []model.Option{
			{OptionNumber: 1, OptionString: "New project", Function:asks.NewProjectsAsks},
			{OptionNumber: 2, OptionString: "Delete project", Function: deleteProject},
			{OptionNumber: 3, OptionString: "Open project", Function: asks.OpenCodeProjectAsks},
			{OptionNumber: 4, OptionString: "Exit", Function: func() bool {return false}},
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

func deleteProject() bool {
	fmt.Println("Deleting project...")
	return true
}
