package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetData() (string, error) {
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
	io.Copy(os.Stdout, resp.Body)
	// // body, err :=
	return "", nil
}
