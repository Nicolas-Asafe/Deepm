package main

import (
	"fmt"
	"os"
	"path/filepath"
	"pc/View/Terminal"
)

func main(){
	if !FolderExists(`C:\Dev\Projects`){
		initializeFolders()
	}
	view.Menu()
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
func initializeFolders() {
	basePath := `C:\Dev\Projects`

	folders := []string{
		"JavaScriptProjects",
		"JavaProjects",
		"GoLangProjects",
	}

	for _, folder := range folders {
		path := filepath.Join(basePath, folder)

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("[x] Falha ao criar %s: %v\n", folder, err)
			continue
		}

		fmt.Printf("[+] Added %s folder\n", folder)
	}

	fmt.Println("[✓] Folders initializers successfully")
}