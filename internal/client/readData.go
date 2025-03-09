package client

import (
	"encoding/json"
	"fmt"

	"github.com/sar0868/currency_converter/internal/models"
)

func GetData(body []byte) (models.CurrenciesData, error) {
	var data models.CurrenciesData
	err := json.Unmarshal(body, &data)
	if err != nil {
		return data, fmt.Errorf("error decode: %w", err)
	}
	rub := models.Data{
		NumCode:  "643",
		CharCode: "RUB",
		Nominal:  1,
		Value:    1,
		Name:     "Российский рубль",
	}
	data.Valute["RUB"] = rub
	return data, nil
}
