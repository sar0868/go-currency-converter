package main

import (
	"github.com/sar0868/currency_converter/internal/app"
	"github.com/sar0868/currency_converter/internal/config"
)

func main() {
	config.Init()
	app.Start()
}
