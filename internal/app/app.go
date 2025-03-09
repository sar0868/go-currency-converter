package app

import (
	"fmt"

	"github.com/sar0868/currency_converter/internal/config"
	"github.com/sar0868/currency_converter/internal/info"
)

func Start() {
	config.Init()
	fmt.Println("run currency converter")
	info.InputInfo()
}
