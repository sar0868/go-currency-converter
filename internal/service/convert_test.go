package service

import (
	"testing"

	"github.com/sar0868/currency_converter/internal/models"
)

func TestConverter(t *testing.T) {
	tests := []struct {
		name    string
		args    models.InputData
		data    models.CurrenciesData
		want    float64
		wantErr bool
	}{
		{
			name: "Converted 1 EUR(2) USD(2): expected 1",
			args: models.InputData{
				Count:   1,
				Valuta1: models.EUR,
				Valuta2: models.USD,
			},
			data:    models.CurrenciesData{},
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Converter(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Converter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Converter() = %v, want %v", got, tt.want)
			}
		})
	}
}
