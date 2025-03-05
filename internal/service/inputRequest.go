package service

import (
	"fmt"

	"github.com/choria-io/go-validator/enum"
	"github.com/sar0868/currency_converter/internal/models"
)

func Input() (models.InputData, error) {
	var input models.InputData

	fmt.Println("input count curr1 curr2")
	var count float64
	var exchanged string
	var received string

	fmt.Scan(&count)
	fmt.Scan(&exchanged)
	fmt.Scan(&received)

	if count <= 0 {
		return input, fmt.Errorf("can't be less than 0: ")
	}
	valid, err := enum.ValidateSlice([]string{exchanged, received}, models.ValidList)
	if err != nil {
		return input, err
	}
	if !valid {
		return input, fmt.Errorf("error input name currency")
	}
	input = models.InputData{
		Count:     count,
		Exchanged: models.ValName(exchanged),
		Received:  models.ValName(received),
	}
	return input, nil
}
