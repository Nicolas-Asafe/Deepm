package asks

import (
	"fmt"
	"io/fs"
	"path/filepath"
	enum "pc/Enum"
	"pc/Kernel/project-manager"
)

func ListProjectsAsks() bool {
	var langOption enum.Lang
	fmt.Println("\nSelect language:")
	fmt.Println("[0] JavaScript")
	fmt.Println("[1] GoLang")
	fmt.Println("[2] Java")
	fmt.Println("[3] Quit")
	fmt.Print("> ")
	fmt.Scanln(&langOption)
    if langOption == 3 {return false}
	projects,ret := project.ListProjects(langOption)
	if !ret{
		println("Error for list projects")
		return true
	} 
	fmt.Println("===================================================")
	fmt.Println(langOption,"PROJECTS")
	for _,project := range projects{
		fmt.Println("===================================================")
		print(project.Name())
		print("- @SIZE=(")
		print(getDirSize(project.Name(),langOption))
		print(")")
		fmt.Println("")
		fmt.Println("===================================================")
	}
	return true
}
func getDirSize(name string,lang enum.Lang) (int64, error) {
	var total int64
	root:=fmt.Sprintf("C:\\Dev\\Projects\\%sProjects\\%s",lang,name)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			total += info.Size()
		}
		return nil
	})

	return total, err
}