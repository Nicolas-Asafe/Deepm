package files

import (
	"fmt"
	"os"
	model "pc/Model"
)

func CreateFile(
	p model.Project,path string,
) {
    example := p.LangProject.MainCodeExample()
	err := os.WriteFile(example[0],[]byte(example[1]),0644)
	if err != nil{
		fmt.Println("Error for create file")
		return
	}
	fmt.Println("File created successfully")
}