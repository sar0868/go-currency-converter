package service

import (
	"math"

	"github.com/sar0868/currency_converter/internal/models"
)

func Converter(inputData models.InputData, data models.CurrenciesData) (float64, error) {
	var result float64
	curVal1 := ValueValute(data, inputData.Exchanged)
	value1 := curVal1 * inputData.Count
	curVal2 := ValueValute(data, inputData.Received)
	result = math.Round((value1/curVal2)*100) / 100
	return result, nil
}

func ValueValute(data models.CurrenciesData, valute models.ValName) float64 {
	return data.Valute[string(valute)].Value / float64(data.Valute[string(valute)].Nominal)
}
