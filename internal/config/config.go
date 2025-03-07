package config

import (
	"log"

	"github.com/sar0868/currency_converter/internal/client"
	"github.com/sar0868/currency_converter/internal/models"
)

var Data models.CurrenciesData

func Init() {
	bytes, errClient := client.Client()
	if errClient != nil {
		log.Fatal(errClient)
	}
	var errParse error
	Data, errParse = client.GetData(bytes)
	if errParse != nil {
		log.Fatal(errParse)
	}
}
