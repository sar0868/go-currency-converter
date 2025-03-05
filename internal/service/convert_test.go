package service

import (
	"testing"

	"github.com/sar0868/currency_converter/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	tests := []struct {
		name string
		args models.InputData
		data models.CurrenciesData
		want float64
	}{
		{
			name: "Converted 1 EUR(2) USD(2): expected 1",
			args: models.InputData{
				Count:     1,
				Exchanged: models.EUR,
				Received:  models.USD,
			},
			data: models.CurrenciesData{
				Valute: map[string]models.Data{
					"EUR": {
						Nominal: 1,
						Value:   2,
					},
					"USD": {
						Nominal: 1,
						Value:   2,
					},
				},
			},
			want: 1,
		},
		{
			name: "Converted 1 EUR(3) USD(2): expected 1.5",
			args: models.InputData{
				Count:     1,
				Exchanged: models.EUR,
				Received:  models.USD,
			},
			data: models.CurrenciesData{
				Valute: map[string]models.Data{
					"EUR": {
						Nominal: 1,
						Value:   3,
					},
					"USD": {
						Nominal: 1,
						Value:   2,
					},
				},
			},
			want: 1.5,
		},
		{
			name: "Converted 3 EUR(nominal 10, value 4), val3) USD(2): expected 0.6",
			args: models.InputData{
				Count:     3,
				Exchanged: models.EUR,
				Received:  models.USD,
			},
			data: models.CurrenciesData{
				Valute: map[string]models.Data{
					"EUR": {
						Nominal: 10,
						Value:   4,
					},
					"USD": {
						Nominal: 1,
						Value:   2,
					},
				},
			},
			want: 0.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Converter(tt.args, tt.data)
			assert.Equal(t, tt.want, result)
		})
	}
}
