package models

type Data struct {
	NumCode  string  `xml:"NumCode" json:"numCode"`
	CharCode string  `xml:"CharCode" json:"charCode"`
	Nominal  int     `xml:"Nominal" json:"nominal"`
	Name     string  `xml:"Name" json:"name"`
	Value    float64 `xml:"Value" json:"value"`
}

type CurrenciesData struct {
	Valute map[string]Data `xml:"Valute" json:"valute"`
}
