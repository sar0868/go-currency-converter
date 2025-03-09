package models

type Data struct {
	NumCode  string  `json:"numCode"`
	CharCode string  `json:"charCode"`
	Nominal  int     `json:"nominal"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
}

type CurrenciesData struct {
	Valute map[string]Data `json:"valute"`
}
