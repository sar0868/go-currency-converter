package service

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/models"
)

func Converter(inputData models.InputData) (float64, error) {
	fmt.Println(inputData)
	return 0, nil
}
