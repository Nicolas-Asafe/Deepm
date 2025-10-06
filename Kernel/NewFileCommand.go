package kernel

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/files"
	model "pc/Model"
)


func NewFileCommand(
	NameProject string,
	LangProject enum.Lang,
	Version float64,
	Author string,
) bool {
	//Initialize kernel func
	p:= model.Project{NameProject: NameProject,LangProject: LangProject,Version: Version,Author: Author}
	p.SetDateCreationForNow()
	k:= NewFuncKernel(&p)
	//Verify if project exists and return a boolean
	res:=files.VerifyProjectExists(k.GetProjectPath())
	if !res {return false}
	//Create file
	files.CreateFile(p.LangProject.MainCodeExample()[0],p.LangProject.MainCodeExample()[1],k.GetProjectPath())

	content := fmt.Sprintf(`@echo off
echo Name project: "%s"
echo Date creation: "%s"
echo Language project: "%s"
echo Version: "%f"
echo Author: "%s"`, NameProject, p.DateCreation, LangProject, Version, Author)


	files.CreateFile("info.bat",content,k.GetProjectPath())
	return true
}