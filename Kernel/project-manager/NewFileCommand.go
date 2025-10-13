package project

import (
	"fmt"
	"os"
	enum "pc/Enum"
	kernel "pc/Kernel"
	"pc/Kernel/files-manager"
	model "pc/Model"
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
	k := kernel.NewFuncKernel(&p)

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

	CreateInfoFile(p,path)
	return true
}

func CreateInfoFile(p model.Project, path string){
		content := fmt.Sprintf(`@echo off
echo Name project: %s
echo Date creation: %s
echo Language project: %s
echo Version: %f
echo Author: %s`,
		p.NameProject, p.DateCreation, p.LangProject, p.Version, p.Author)

	files.CreateFile("info.bat", content, path)
}