package service

import (
	"testing"

	"github.com/sar0868/currency_converter/internal/models"
)

func TestConverter(t *testing.T) {
	type args struct {
		valuetes []models.CurrenciesData
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Converter(tt.args.valuetes)
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
