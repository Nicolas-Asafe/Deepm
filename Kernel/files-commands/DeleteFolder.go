package files

import (
	"fmt"
	"os"
)

func DeleteFolder(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error for delete folder:", err)
		return false
	} else {
		fmt.Println("Folder deleted successfully")
	}
	return true
}
