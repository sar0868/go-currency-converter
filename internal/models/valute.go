package models

type Valute struct {
	ID      string  `json:"id"`
	NumCode string  `json:"numCode"`
	Nominal int     `json:"nominal"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
}

type CurrenciesData struct {
	Date   string `json:"date"`
	Values map[string]Valute
}
