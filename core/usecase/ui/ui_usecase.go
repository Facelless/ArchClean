package ui

import (
	"faceclean/core/usecase"
	"fmt"

	"github.com/manifoldco/promptui"
)

func Archtype() {
	menu := []string{
		"Limpar arquivos",
	}
	prompt := promptui.Select{
		Label: "Escolha uma opção",
		Items: menu,
	}

	index, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)

	switch index {
	case 0:
		usecase.ClearArchives()
	}
}
