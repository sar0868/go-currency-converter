package client

import (
	"fmt"
	"io"
	"net/http"
)

func Client() ([]byte, error) {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	httpClient := &http.Client{}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error HTTP-response: %v", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read: %w", err)
	}
	return body, nil
}
