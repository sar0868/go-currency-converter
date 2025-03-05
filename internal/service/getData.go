package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sar0868/currency_converter/internal/models"
)

func GetData() (models.CurrenciesData, error) {
	var data models.CurrenciesData

	// url := "https://www.cbr-xml-daily.ru/daily_utf8.xml"
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	httpClient := &http.Client{}
	resp, err := httpClient.Get(url)
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
	return data, nil
}
