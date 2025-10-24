package project

import (
	"fmt"
	enum "pc/Enum"
	"pc/Kernel/files-commands"
)

func DeleteProjectCommand(
	NameProject string,
	LangOption enum.Lang,
) bool{

	path:=fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\%s", LangOption.String(),NameProject)

	if !files.DeleteFolder(path) { return false }
	fmt.Println("Project deleted successfully")
	return true
}
