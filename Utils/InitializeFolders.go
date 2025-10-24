package Utils

import (
	"fmt"
	"os"
	"path/filepath"
	enum "pc/Enum"
)

func InitializeFolders() {
	basePath := `C:\Dev\Projects`

	folders := enum.GetAllLanguagesStrings()
	for _, folder := range folders {
		path := filepath.Join(basePath, folder)

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("[x] Error for create %s: %v\n", folder, err)
			continue
		}

		fmt.Printf("[+] Added %s folder\n", folder)
	}

	fmt.Println("[✓] Folders initializers successfully")
}