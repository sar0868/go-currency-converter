package service

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/models"
)

func Converter(valuetes []models.CurrenciesData) (float64, error) {
	fmt.Println(valuetes)
	return 0, nil
}
