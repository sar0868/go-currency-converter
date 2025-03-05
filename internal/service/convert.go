package service

import (
	"math"

	"github.com/sar0868/currency_converter/internal/models"
)

func Converter(inputData models.InputData, data models.CurrenciesData) (float64, error) {
	rub := models.Data{
		NumCode:  "643",
		CharCode: "RUB",
		Nominal:  1,
		Value:    1,
		Name:     "Российский рубль",
	}
	data.Valute["RUB"] = rub
	var result float64
	curVal1 := ValueValute(data, inputData.Exchanged)
	value1 := curVal1 * inputData.Count
	curVal2 := ValueValute(data, inputData.Received)
	result = math.Round((value1/curVal2)*100) / 100
	return result, nil
}

func ValueValute(data models.CurrenciesData, valute models.ValName) float64 {
	return math.Round(data.Valute[string(valute)].Value*100/float64(data.Valute[string(valute)].Nominal)) / 100
}
