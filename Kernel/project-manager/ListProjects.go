package project

import (
	"fmt"
	"os"
	enum "pc/Enum"
	"pc/Kernel/files-commands"
)

func ListProjects(langOption enum.Lang) ([]os.DirEntry,bool){
	var langProject enum.Lang
	switch langOption {
	case 0:
		langProject = enum.JavaScript
	case 1:
		langProject = enum.GoLang
	case 2:
		langProject = enum.Java
	default:
		fmt.Println("Invalid language option.")
		return nil,true
	}
	path := fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\", langProject)
	listProjects,ret := files.ListDir(path)
	if !ret{
		return nil,ret
	}
	if len(listProjects) != 0{
		return listProjects,true
	}
	return nil,true
}