package main

import (
	"os"
	"pc/View/Terminal"
	Utils "pc/utils"
)

func main(){
	if !FolderExists(`C:\Dev\Projects`){
		Utils.InitializeFolders()
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