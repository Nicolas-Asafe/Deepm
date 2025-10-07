package kernel

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/files"
	model "pc/Model"
	"os"
)


func NewFileCommand(
	NameProject string,
	LangProject enum.Lang,
	Version float64,
	Author string,
) bool {
	
	p := model.Project{
		NameProject: NameProject,
		LangProject: LangProject,
		Version:     Version,
		Author:      Author,
	}
	p.SetDateCreationForNow()
	k := NewFuncKernel(&p)

	path := k.GetProjectPath()
	fmt.Println("Project path:", path)


	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar pasta do projeto:", err)
		return false
	}


	files.CreateFile(
		p.LangProject.MainCodeExample()[0],
		p.LangProject.MainCodeExample()[1],
		path,
	)


	content := fmt.Sprintf(`@echo off
echo Name project: "%s"
echo Date creation: "%s"
echo Language project: "%s"
echo Version: "%f"
echo Author: "%s"`,
		NameProject, p.DateCreation, LangProject, Version, Author)

	files.CreateFile("info.bat", content, path)

	return true
}
