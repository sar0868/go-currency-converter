package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sar0868/currency_converter/internal/models"
)

func GetData() (models.CurrenciesData, error) {
	var data models.CurrenciesData
	// data := make(map[string]models.Valute)

	// url := "https://www.cbr-xml-daily.ru/daily_utf8.xml"
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url) //nolint: noctx
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("error get data")
	}
	errDecode := json.NewDecoder(resp.Body).Decode(&data)
	// errDecode := xml.NewDecoder(resp.Body).Decode(&data)
	if errDecode != nil {
		return data, fmt.Errorf("error decode: %w", errDecode)
	}

	fmt.Println(data)
	fmt.Println(data.Valute["AED"].Value)
	fmt.Println(data.Valute["AED"].Name)
	fmt.Println(data.Valute["AED"].NumCode)
	return data, nil
}
