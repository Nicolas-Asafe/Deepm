package files

import (
	"fmt"
	"os"
)

func VerifyProjectExists(path string) bool{
	if _,err:= os.Stat(path); os.IsNotExist(err){
		fmt.Println("Project not exists")
		return true
	}else{
		fmt.Println("Project already exists")
		return false
	}
}