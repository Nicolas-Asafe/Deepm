package files

import (
	"fmt"
	"os"
)

func CreateFile(
	name string,content string,path string,
) {
	err := os.WriteFile(name,[]byte(content),0644)
	if err != nil{
		fmt.Println("Error for create file")
		return
	}
	fmt.Println("File created successfully")
}