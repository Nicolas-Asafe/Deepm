package project

import (
	"fmt"
	enum "pc/Enum"
	kernel "pc/Kernel"
	"pc/Kernel/files-commands"
	model "pc/Model"
)

func DeleteProjectCommand(
	NameProject string,
	LangOption int,
) bool{
	p:=model.Project{
		NameProject: NameProject,
		LangProject: 3,
	}

	switch LangOption{
	case 0:
		p.LangProject = enum.JavaScript
	case 1:
		p.LangProject = enum.GoLang
	case 2:
		p.LangProject = enum.Java
	default:
		fmt.Println("Invalid language option.")
		return true
	}

	k:=kernel.NewFuncKernel(&p)
	path:=k.GetProjectPath()

	if !files.DeleteFolder(path) { return false }
	fmt.Println("Project deleted successfully")
	return true
}
