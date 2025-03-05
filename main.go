package main

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/info"
	"github.com/sar0868/currency_converter/internal/service"
)

func main() {
	service.GetData()

	fmt.Println("run currency converter")
	info.InputInfo()
}
