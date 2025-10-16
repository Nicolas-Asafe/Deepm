package project

import (
	"fmt"
	enum "pc/Enum"
	kernel "pc/Kernel"
	files "pc/Kernel/files-commands"
	model "pc/Model"
)

func NewProjectCommand(
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

	if !files.CreateFolder(path) {
		return false
	}

	if !CreateProjectFiles(p, path) {
		return false
	}

	return true
}

func CreateProjectFiles(p model.Project, path string) bool {
	files.CreateFile(
		p.LangProject.MainCodeExample()[0],
		p.LangProject.MainCodeExample()[1],
		path,
	)
	CreateInfoFile(p, path)
	return true
}

func CreateInfoFile(p model.Project, path string) bool {
	content := fmt.Sprintf(`@echo off
echo Name project: %s
echo Date creation: %s
echo Language project: %s
echo Version: %f
echo Author: %s`,
		p.NameProject, p.DateCreation, p.LangProject, p.Version, p.Author)

	files.CreateFile("info.bat", content, path)
	return true
}
