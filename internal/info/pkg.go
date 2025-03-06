package info

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/models"
	"github.com/sar0868/currency_converter/internal/service"
)

func InputInfo() {
	Menu()
	Actions(ChoiceAction())
}

func Show() {}

func Menu() {
	fmt.Println(`Расчет обмена валют:
1. Расчет обмена валют
2. Список обозначения валют
3. Порядок расчета обмена валют
4. Выход`)
}

func ChoiceAction() string {
	var choice string
	fmt.Scan(&choice)
	return choice
}

func Actions(action string) {
	switch action {
	case "1":
		_, err := service.Input()
		if err != nil {
			fmt.Printf("Error input request: %v\n", err)
		}
	case "2":
		fmt.Println(models.ValidList)
	case "3":
		fmt.Println("instruction")
	case "4":
		fmt.Println("exit")
	}
}
