package main

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/client"
	"github.com/sar0868/currency_converter/internal/info"
	"github.com/sar0868/currency_converter/internal/service"
)

func main() {
	bytes, errClient := client.Client()
	if errClient != nil {
		fmt.Println(errClient)
	}
	_, errParse := service.GetData(bytes)
	if errParse != nil {
		fmt.Println(errParse)
	}

	fmt.Println("run currency converter")
	info.InputInfo()
}
