package view

import (
	"fmt"
	model "pc/Model"
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
			{OptionNumber: 2, OptionString: "Delete project", Function: asks.DeleteProjectAsks},
			{OptionNumber: 3, OptionString: "Open project", Function: asks.OpenCodeProjectAsks},
			{OptionNumber: 4, OptionString: "List projects", Function:asks.ListProjectsAsks},//
			{OptionNumber: 5, OptionString: "Exit", Function: func() bool {return false}},
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