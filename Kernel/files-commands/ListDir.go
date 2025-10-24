package files

import (
	"fmt"
	"os"
)

func ListDir(path string) ([]os.DirEntry,bool) {
	entries, err := os.ReadDir(path)
	var dirs []os.DirEntry

	if err !=nil{
		fmt.Println("Error for list dirs: ",err)
		return nil,false
	}

	for _,entrie := range entries{
		if entrie.IsDir(){
			dirs = append(
				dirs,
				entrie,
			)
		}
	}
	return dirs,true
}