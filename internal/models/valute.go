package models

type Data struct {
	ID      string
	NumCode string
	Nominal int
	Name    string
	Value   float64
	// ID      string  `json:"id"`
	// NumCode string  `json:"numCode"`
	// Nominal int     `json:"nominal"`
	// Name    string  `json:"name"`
	// Value   float64 `json:"value"`
}

type CurrenciesData struct {
	Valute map[string]Data
	// Date   string `json:"date"`
}
