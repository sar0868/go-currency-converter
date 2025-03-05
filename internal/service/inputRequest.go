package service

import (
	"fmt"

	"github.com/choria-io/go-validator/enum"
	"github.com/sar0868/currency_converter/internal/models"
)

func Input() (models.InputData, error) {
	// var inp string
	var input models.InputData
	fmt.Println(input)

	fmt.Println("input count curr1 curr2")
	var count float64
	var exchanged string
	var received string
	// fmt.Scan(&inp)
	// fmt.Println(inp)
	// inpArr := strings.Split(inp, " ")
	// fmt.Println(inpArr)

	// if len(inpArr) != 3 {
	// 	return input, fmt.Errorf("incorrect input data")
	// }
	// count, errCnt := strconv.ParseFloat(inpArr[0], 64)
	// if errCnt != nil {
	// 	return input, errCnt
	// }
	// exchanged := inpArr[1]
	// received := inpArr[2]
	fmt.Scan(&count)
	fmt.Scan(&exchanged)
	fmt.Scan(&received)

	if count <= 0 {
		return input, fmt.Errorf("can't be less than 0: ")
	}
	valid, err := enum.ValidateSlice(models.ValidList, []string{exchanged, received})
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
	fmt.Println(input)
	return input, nil
}
