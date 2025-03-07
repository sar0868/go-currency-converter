package service

import (
	"testing"

	"github.com/sar0868/currency_converter/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_parseInputData(t *testing.T) {
	type args struct {
		count     float64
		exchanged string
		received  string
	}
	tests := []struct {
		name    string
		args    args
		want    models.InputData
		wantErr bool
	}{
		{
			name: "correct input 1 RUB USD",
			args: args{
				count:     1,
				exchanged: "RUB",
				received:  "USD",
			},
			want: models.InputData{
				Count:     1,
				Exchanged: models.RUB,
				Received:  models.USD,
			},
			wantErr: false,
		},
		{
			name: "incorrect input -1 RUB USD",
			args: args{
				count:     -1,
				exchanged: "RUB",
				received:  "USD",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 1 rub USD",
			args: args{
				count:     1,
				exchanged: "rub",
				received:  "USD",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 1 RUB US",
			args: args{
				count:     1,
				exchanged: "RUB",
				received:  "US",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 1 RUB",
			args: args{
				count:     1,
				exchanged: "RUB",
				received:  "",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 1 RUBUSD",
			args: args{
				count:     1,
				exchanged: "RUBUSD",
				received:  "",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 0 RUB USD",
			args: args{
				count:     0,
				exchanged: "RUB",
				received:  "USD",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
		{
			name: "incorrect input 1 RUB MMM",
			args: args{
				count:     1,
				exchanged: "RUB",
				received:  "MMM",
			},
			want: models.InputData{
				Count: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInputData(tt.args.count, tt.args.exchanged, tt.args.received)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
