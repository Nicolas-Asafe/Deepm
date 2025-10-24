package asks

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/project-manager"
)

func DeleteProjectAsks() bool {
	var NameProject string
	var LangOption enum.Lang

	fmt.Print("Enter the name project: ")
	fmt.Scanln(&NameProject)
	if NameProject == "" {
		fmt.Println("Please enter the name project")
	}
	fmt.Println("\nSelect language:")
	fmt.Println("[0] JavaScript")
	fmt.Println("[1] GoLang")
	fmt.Println("[2] Java")
	fmt.Println("[3] Quit")
	fmt.Print("> ")
	fmt.Scanln(&LangOption)
    if LangOption == 3 {return false}

	project.DeleteProjectCommand(NameProject,LangOption)

	return true
}