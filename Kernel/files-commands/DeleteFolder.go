package files

import (
	"fmt"
	"os"
)

func DeleteFolder(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error for delete folder:", err)
	} else {
		fmt.Println("Folder deleted successfully")
	}
}
