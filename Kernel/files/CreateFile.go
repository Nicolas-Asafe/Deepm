package files

import (
	"fmt"
	"os"
	"path/filepath"
)


func CreateFile(name, content, path string) {
	fullPath := filepath.Join(path, name) 

	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created successfully:", fullPath)
}
