package service

import (
	"github.com/sar0868/currency_converter/internal/models"
)

func Converter(inputData models.InputData, data models.CurrenciesData) float64 {
	var result float64
	curVal1 := ValueValute(data, inputData.Exchanged)
	value1 := curVal1 * inputData.Count
	curVal2 := ValueValute(data, inputData.Received)
	result = value1 / curVal2
	return result
}

func ValueValute(data models.CurrenciesData, valute models.ValName) float64 {
	return data.Valute[string(valute)].Value / float64(data.Valute[string(valute)].Nominal)
}
