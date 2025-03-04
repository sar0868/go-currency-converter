package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sar0868/currency_converter/internal/models"
)

func GetData() (string, error) {
	var data models.CurrenciesData
	// data := make(map[string]models.Valute)
	// url := "https://www.cbr.ru/scripts/XML_daily.asp"
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url) //nolint: noctx
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error get data")
	}
	errDecode := json.NewDecoder(resp.Body).Decode(&data)
	if errDecode != nil {
		return "", fmt.Errorf("error decode: %w", errDecode)
	}
	// for val := range data {
	// 	for v := range val {
	// 		fmt.Println(v)
	// 	}

	// }
	fmt.Println(data)
	// // body, err :=
	return "", nil
}
