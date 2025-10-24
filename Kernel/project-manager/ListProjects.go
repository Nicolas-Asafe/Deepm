package project

import (
	"fmt"
	"os"
	enum "pc/Enum"
	"pc/Kernel/files-commands"
)

func ListProjects(langOption enum.Lang) ([]os.DirEntry,bool){
	var langProject enum.Lang = langOption

	path := fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\", langProject.String())
	listProjects,ret := files.ListDir(path)
	if !ret{
		return nil,ret
	}
	if len(listProjects) != 0{
		return listProjects,true
	}
	return nil,true
}