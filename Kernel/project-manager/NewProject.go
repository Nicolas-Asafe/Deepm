package project

import (
	"fmt"
	enum "pc/Enum"
	files "pc/Kernel/files-commands"
	model "pc/Model"
)

func NewProjectCommand(
	NameProject string,
	LangOption enum.Lang,
	Version float64,
	Author string,
) bool {
	var LangProject enum.Lang = LangOption
	
	path := fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\%s", LangProject.String(),NameProject)
	fmt.Println("Project path:", path)

	if !files.CreateFolder(path) {
		return false
	}
	p:=model.Project{
		NameProject: NameProject,
		LangString: LangOption.String(),
		Version: Version,
		Author: Author,
	}

	return CreateProjectFiles(p,LangOption,path) 
}

func CreateProjectFiles(p model.Project,LangOption enum.Lang, path string) bool {
	files.CreateFile(
		LangOption.MainCodeExample()[0],
		LangOption.MainCodeExample()[1],
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
		p.NameProject, p.DateCreation, p.LangString, p.Version, p.Author)

	files.CreateFile("info.bat", content, path)
	return true
}
