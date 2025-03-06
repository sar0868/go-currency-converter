package service

import (
	"testing"

	"github.com/sar0868/currency_converter/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	type args struct {
		bodyStr string
	}
	tests := []struct {
		name    string
		args    args
		want    models.CurrenciesData
		wantErr bool
	}{
		{
			name: "correct parse []byte",
			args: args{
				bodyStr: `{"valute":{"RUB":{"numCode":"643",
						"charCode":"RUB","nominal":1,"name":"Российский рубль",
						"value":1},"USD":{"numCode":"840","charCode":"USD",
						"nominal":1,"name":"Доллар США","value":89.7878}}}`,
			},
			want: models.CurrenciesData{
				Valute: map[string]models.Data{
					"RUB": {
						NumCode:  "643",
						CharCode: "RUB",
						Nominal:  1,
						Value:    1,
						Name:     "Российский рубль",
					},
					"USD": {
						NumCode:  "840",
						CharCode: "USD",
						Nominal:  1,
						Value:    89.7878,
						Name:     "Доллар США",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(tt.args.bodyStr)
			got, err := GetData(body)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
