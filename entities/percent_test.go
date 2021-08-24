package entities

import (
	"math/big"
	"testing"

	"github.com/miraclesu/uniswap-sdk-go/number"
)

// nolint:dupl
func TestPercent_ToSignificant(t *testing.T) {
	type fields struct {
		num, deno *big.Int
	}
	type args struct {
		significantDigits uint
		opt               []number.Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"returns the value scaled by 100",
			fields{big.NewInt(154), big.NewInt(10000)},
			args{significantDigits: 3},
			"1.54",
		},
	}
	for _, tt := range tests {
		got := NewPercent(tt.fields.num, tt.fields.deno).ToSignificant(tt.args.significantDigits, tt.args.opt...)
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got != want {
				t.Errorf("ToSignificant() = %v, want %v", got, want)
			}
		})
	}
}

// nolint: dupl
func TestPercent_ToFixed(t *testing.T) {
	type fields struct {
		num, deno *big.Int
	}
	type args struct {
		decimalPlaces uint
		opt           []number.Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"returns the value scaled by 100",
			fields{big.NewInt(154), big.NewInt(10000)},
			args{decimalPlaces: 2},
			"1.54",
		},
	}
	for _, tt := range tests {
		got := NewPercent(tt.fields.num, tt.fields.deno).ToFixed(tt.args.decimalPlaces, tt.args.opt...)
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got != want {
				t.Errorf("ToFixed() = %v, want %v", got, want)
			}
		})
	}
}
