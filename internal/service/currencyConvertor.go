package service

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/config"
	"github.com/sar0868/currency_converter/internal/models"
)

func CurrencyConvertor() (float64, models.InputData, error) {
	var result float64
	inputData, err := Input()
	if err != nil {
		return result, inputData, fmt.Errorf("error input request: %w", err)
	}
	result, errConverted := Converter(inputData, config.Data)
	if errConverted != nil {
		return result, inputData, fmt.Errorf("error converter: %w", errConverted)
	}
	return result, inputData, nil
}
