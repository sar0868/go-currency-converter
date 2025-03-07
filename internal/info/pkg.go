package info

import (
	"fmt"
	"os"

	"github.com/sar0868/currency_converter/internal/config"
	"github.com/sar0868/currency_converter/internal/models"
	"github.com/sar0868/currency_converter/internal/service"
)

func InputInfo() {
	for {
		fmt.Println("\n======================")
		Menu()
		Actions(choiceAction())
	}
}

func Show(result float64, req models.InputData) {
	currencyName := config.Data.Valute[string(req.Received)].Name
	fmt.Println("--------------------------")
	fmt.Printf("результат %.4f %s (%v)\n", result, req.Received, currencyName)
}

func Menu() {
	fmt.Println(`Расчет обмена валют:
1. Расчет обмена валют
2. Список обозначения валют
3. Порядок расчета обмена валют
4. Выход`)
}

func choiceAction() string {
	var choice string
	fmt.Print("Введите Ваш выбор: ")
	fmt.Scan(&choice)
	return choice
}

func Actions(action string) {
	switch action {
	case "1":
		result, inputData, err := service.CurrencyConvertor()
		if err != nil {
			fmt.Println(err)
		}
		Show(result, inputData)
	case "2":
		printListCurrencies()
	case "3":
		msg := `Для расчета курса валют необходимо ввести в консоле
count[количество обменеваемой валюты] exchenged[название 
обменеваемой валюты из списка валют] received[название 
получаемой валюты из списка валют]`
		fmt.Printf("\n%s\n", msg)
	case "4":
		fmt.Println("exit")
		os.Exit(0)
	}
}

func printListCurrencies() {
	fmt.Println("\nСписок обозначений валют:")
	for i := 0; i < len(models.ValidList); {
		for j := 0; j < 8; j++ {
			if (i + j) >= len(models.ValidList) {
				continue
			}
			fmt.Printf("%s ", models.ValidList[i+j])
		}
		fmt.Println()
		i += 8
	}
}
