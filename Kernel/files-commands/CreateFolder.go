package files

import (
	"fmt"
	"os"
)

func CreateFolder(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar pasta do projeto:", err)
		return false
	}
	return true
}