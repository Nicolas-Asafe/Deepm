package View

import (
	"fmt"
	"pc/Model"
)

func RenderOptions(options []model.Option) (bool, error) {
	for _, p := range options {
		fmt.Printf("[%d] - %s\n", p.OptionNumber, p.OptionString)
	}

	fmt.Print("> ")
	var res int
	_, err := fmt.Scanln(&res)
	if err != nil {
		return true, fmt.Errorf("invalid input")
	}

	for _, p := range options {
		if res == p.OptionNumber {
			return p.Function(), nil
		}
	}

	fmt.Println("Invalid option, try again.")
	return true, nil
}
