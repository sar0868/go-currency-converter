package info

import "fmt"

func InputInfo() {
	Menu()
	fmt.Println(ChoiceAction())
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
		fmt.Println("Расчет")
	case "2":
		fmt.Println("print slice currency")
	case "3":
		fmt.Println("instruction")
	case "4":
		fmt.Println("exit")
	}
}
