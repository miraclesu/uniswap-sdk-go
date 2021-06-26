package number

import (
	"testing"

	"github.com/shopspring/decimal"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

func mustNewFromString(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}

	return d
}

func TestDecimalFormat(t *testing.T) {
	t.Parallel()

	type args struct {
		d    decimal.Decimal
		opts *Options
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				d:    decimal.NewFromInt(0),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "1",
		},
		{
			args: args{
				d:    decimal.NewFromInt(-1),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "-1",
		},
		{
			args: args{
				d:    decimal.NewFromFloat(123.456),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "123.456",
		},
		{
			args: args{
				d:    decimal.NewFromFloat(123.456),
				opts: New(WithFractionGroupSeparator(' '), WithDecimalPlaces(3)),
			},
			want: "123.456",
		},
		{
			args: args{
				d:    decimal.NewFromInt(0),
				opts: New(WithFractionGroupSeparator(' '), WithDecimalPlaces(1)),
			},
			want: "0.0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithDecimalPlaces(2)),
			},
			want: "1.00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(-1),
				opts: New(WithFractionGroupSeparator(' '), WithDecimalPlaces(3)),
			},
			want: "-1.000",
		},
		{
			args: args{
				d:    decimal.NewFromFloat(123.456),
				opts: New(WithFractionGroupSeparator(' '), WithDecimalPlaces(4)),
			},
			want: "123.4560",
		},
		{
			args: args{
				d:    decimal.NewFromFloat(9876.54321),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9,876.54321",
		},
		{
			args: args{
				d:    mustNewFromString("4.0187364e+21"),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "4,018,736,400,000,000,000,000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "999,999,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "99,999,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9,999,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "999,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "99,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9,999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "999,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "99,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9,999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "999,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "99",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "9",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091e+4"),
				opts: New(WithFractionGroupSeparator(' ')),
			},
			want: "76,852.342091",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithDecimalPlaces(2)),
			},
			want: "76 852.34",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' ')),
			},
			want: "76 852.342091",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091087145832640897e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithDecimalPlaces(10)),
			},
			want: "76 852.3420910871",
		},
		{
			args: args{
				d:    mustNewFromString("4.0187364e+21"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5)),
			},
			want: "4 018 736 400 000 000 000 000",
		},
		{
			args: args{
				d:    mustNewFromString("7.685234209108714583264089e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(20)),
			},
			want: "76 852.34209 10871 45832 64089",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091087145832640897e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(21)),
			},
			want: "76 852.34209 10871 45832 64089 7",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091087145832640897e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(25)),
			},
			want: "76 852.34209 10871 45832 64089 70000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(0)),
			},
			want: "999 999 999 999 999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(1)),
			},
			want: "99 999 999 999 999.0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(2)),
			},
			want: "9 999 999 999 999.00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(3)),
			},
			want: "999 999 999 999.000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(4)),
			},
			want: "99 999 999 999.0000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(5)),
			},
			want: "9 999 999 999.00000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(6)),
			},
			want: "999 999 999.00000 0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(7)),
			},
			want: "99 999 999.00000 00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(8)),
			},
			want: "9 999 999.00000 000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(9)),
			},
			want: "999 999.00000 0000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(10)),
			},
			want: "99 999.00000 00000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(11)),
			},
			want: "9 999.00000 00000 0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(12)),
			},
			want: "999.00000 00000 00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(13)),
			},
			want: "99.00000 00000 000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(14)),
			},
			want: "9.00000 00000 0000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(15)),
			},
			want: "1.00000 00000 00000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(14)),
			},
			want: "1.00000 00000 0000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(13)),
			},
			want: "1.00000 00000 000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(12)),
			},
			want: "1.00000 00000 00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(11)),
			},
			want: "1.00000 00000 0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(10)),
			},
			want: "1.00000 00000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(5), WithDecimalPlaces(9)),
			},
			want: "1.00000 0000",
		},
		{
			args: args{
				d:    mustNewFromString("4.0187364e+21"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0)),
			},
			want: "4 018 736 400 000 000 000 000",
		},
		{
			args: args{
				d:    mustNewFromString("7.685234209108714583264089e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(20)),
			},
			want: "76 852.34209108714583264089",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091087145832640897e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(21)),
			},
			want: "76 852.342091087145832640897",
		},
		{
			args: args{
				d:    mustNewFromString("7.6852342091087145832640897e+4"),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(25)),
			},
			want: "76 852.3420910871458326408970000",
		},

		{
			args: args{
				d:    decimal.NewFromInt(999999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(0)),
			},
			want: "999 999 999 999 999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(1)),
			},
			want: "99 999 999 999 999.0",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(2)),
			},
			want: "9 999 999 999 999.00",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(3)),
			},
			want: "999 999 999 999.000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(4)),
			},
			want: "99 999 999 999.0000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(5)),
			},
			want: "9 999 999 999.00000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(6)),
			},
			want: "999 999 999.000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(7)),
			},
			want: "99 999 999.0000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(8)),
			},
			want: "9 999 999.00000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(9)),
			},
			want: "999 999.000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(10)),
			},
			want: "99 999.0000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(11)),
			},
			want: "9 999.00000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(12)),
			},
			want: "999.000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(13)),
			},
			want: "99.0000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(14)),
			},
			want: "9.00000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(15)),
			},
			want: "1.000000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(14)),
			},
			want: "1.00000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(13)),
			},
			want: "1.0000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(12)),
			},
			want: "1.000000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(11)),
			},
			want: "1.00000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(10)),
			},
			want: "1.0000000000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(1),
				opts: New(WithFractionGroupSeparator(' '), WithGroupSeparator(' '), WithFractionGroupSize(0), WithDecimalPlaces(9)),
			},
			want: "1.000000000",
		},
		{
			args: args{
				d:    decimal.NewFromFloat(9876.54321),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,876.54321",
		},
		{
			args: args{
				d:    mustNewFromString("1000037.123456789"),
				opts: New(WithSecondaryGroupSize(2), WithDecimalPlaces(3)),
			},
			want: "10,00,037.123",
		},
		{
			args: args{
				d:    mustNewFromString("4.0187364e+21"),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "4,01,87,36,40,00,00,00,00,00,000",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,99,99,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,99,99,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,99,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,99,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9,999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(999),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "999",
		},
		{
			args: args{
				d:    decimal.NewFromInt(99),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "99",
		},
		{
			args: args{
				d:    decimal.NewFromInt(9),
				opts: New(WithSecondaryGroupSize(2)),
			},
			want: "9",
		},
		{
			args: args{
				d:    mustNewFromString("1.23456000000000000000789e+9"),
				opts: New(WithSecondaryGroupSize(2), WithDecimalSeparator(','), WithGroupSeparator('.'), WithDecimalPlaces(12)),
			},
			want: "1.23.45.60.000,000000000008",
		},
		{
			args: args{
				d:    mustNewFromString("10000000000123456789000000.000000000100000001"),
				opts: New(WithSecondaryGroupSize(2), WithDecimalSeparator(','), WithGroupSeparator('\xA0'), WithDecimalPlaces(10)),
			},
			want: "10000000000123456789000000,0000000001",
		},
	}

	for i, tt := range tests {
		if got := DecimalFormat(tt.args.d, tt.args.opts); got != tt.want {
			t.Errorf("DecimalFormat([%d]{d:[%+v], opts:[%+v]}) got = %v, want %v", i, tt.args.d.String(), tt.args.opts,
				got, tt.want)
		}
	}
}

func TestDecimalRound(t *testing.T) {
	type args struct {
		d    decimal.Decimal
		opts *Options
	}
	tests := []struct {
		args args
		want decimal.Decimal
	}{
		{
			args: args{
				d:    mustNewFromString("-0"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("0.00001000"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.0001"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0011117"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.0011117"),
		},
		{
			args: args{
				d:    mustNewFromString("-7187"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-7187"),
		},
		{
			args: args{
				d:    mustNewFromString("4093"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("4093"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.049635"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.0496"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.05"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.1"),
		},
		{
			args: args{
				d:    mustNewFromString("520216039"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("520216039"),
		},
		{
			args: args{
				d:    mustNewFromString("53937.399"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("53937.3"),
		},
		{
			args: args{
				d:    mustNewFromString("63101619"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("63101619"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00207"),
				opts: New(WithRoundingPrecision(6), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.00207"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.03169086"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.032"),
		},
		{
			args: args{
				d:    mustNewFromString("-3583"),
				opts: New(WithRoundingPrecision(5), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-3583"),
		},
		{
			args: args{
				d:    mustNewFromString("-6.615375"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-6.616"),
		},
		{
			args: args{
				d:    mustNewFromString("7528739"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("7528739"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0002613"),
				opts: New(WithRoundingPrecision(6), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.000262"),
		},
		{
			args: args{
				d:    mustNewFromString("0.00009075"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.00009075"),
		},
		{
			args: args{
				d:    mustNewFromString("-4.4195"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-4.42"),
		},
		{
			args: args{
				d:    mustNewFromString("43759"),
				opts: New(WithRoundingPrecision(9), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("43759"),
		},
		{
			args: args{
				d:    mustNewFromString("336379823"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("336379823"),
		},
		{
			args: args{
				d:    mustNewFromString("310614"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("310614"),
		},
		{
			args: args{
				d:    mustNewFromString("-5446775"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-5446775"),
		},
		{
			args: args{
				d:    mustNewFromString("59.7954405"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("59.7954405"),
		},
		{
			args: args{
				d:    mustNewFromString("47085.84"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("47086"),
		},
		{
			args: args{
				d:    mustNewFromString("-2564"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-2564"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000055732"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.00005573"),
		},
		{
			args: args{
				d:    mustNewFromString("-7"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-7"),
		},
		{
			args: args{
				d:    mustNewFromString("-609"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-609"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00004595"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.00004595"),
		},
		{
			args: args{
				d:    mustNewFromString("-22243"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-22243"),
		},
		{
			args: args{
				d:    mustNewFromString("54693"),
				opts: New(WithRoundingPrecision(5), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("54693"),
		},
		{
			args: args{
				d:    mustNewFromString("3808"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("3808"),
		},
		{
			args: args{
				d:    mustNewFromString("0.00008892"),
				opts: New(WithRoundingPrecision(5), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0.00009"),
		},
		{
			args: args{
				d:    mustNewFromString("0.06922"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0.07"),
		},
		{
			args: args{
				d:    mustNewFromString("-326"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-326"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.003"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("-17.14"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-17.14"),
		},
		{
			args: args{
				d:    mustNewFromString("-3.767032983"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-3.76703298"),
		},
		{
			args: args{
				d:    mustNewFromString("0.00023235"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.0003"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0155175"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.0155175"),
		},
		{
			args: args{
				d:    mustNewFromString("-645406477.5"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-645406477.5"),
		},
		{
			args: args{
				d:    mustNewFromString("4611867124"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("4611867124"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0080252"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.0080252"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.02829833"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("-8"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-8"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00735541"),
				opts: New(WithRoundingPrecision(7), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.0073554"),
		},
		{
			args: args{
				d:    mustNewFromString("-3903"),
				opts: New(WithRoundingPrecision(6), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-3903"),
		},
		{
			args: args{
				d:    mustNewFromString("228"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("228"),
		},
		{
			args: args{
				d:    mustNewFromString("805.2467"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("805.25"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0259"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.026"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.008007359"),
				opts: New(WithRoundingPrecision(5), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.00801"),
		},
		{
			args: args{
				d:    mustNewFromString("-706"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-706"),
		},
		{
			args: args{
				d:    mustNewFromString("36.9109527"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("36.911"),
		},
		{
			args: args{
				d:    mustNewFromString("5040908"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("5040908"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.08"),
				opts: New(WithRoundingPrecision(2), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.08"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.438905"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.4389"),
		},
		{
			args: args{
				d:    mustNewFromString("-8858843.829"),
				opts: New(WithRoundingPrecision(3), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-8858843.829"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000005897937615245886508878530770431196412050562641578155968"),
				opts: New(WithRoundingPrecision(34), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.0000058979376152458865088785307705"),
		},
		{
			args: args{
				d:    mustNewFromString("-6049825281564367887763596943301191584240212075976455"),
				opts: New(WithRoundingPrecision(53), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-6049825281564367887763596943301191584240212075976455"),
		},
		{
			args: args{
				d:    mustNewFromString("-64680661822322715719008107701612.74131236713131820297696442216284615573809"),
				opts: New(WithRoundingPrecision(27), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-64680661822322715719008107701612.741312367131318202976964422"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0008577428383491818493352557962418994540276420616048890965876345513"),
				opts: New(WithRoundingPrecision(44), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0.00085774283834918184933525579624189945402764"),
		},
		{
			args: args{
				d:    mustNewFromString("-127706837731025454069338274697755478243.226555768723254468591"),
				opts: New(WithRoundingPrecision(21), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-127706837731025454069338274697755478243.226555768723254468591"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000000000000000006"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00000000000000000000000000000000000000000000030032464295099044566372323"),
				opts: New(WithRoundingPrecision(46), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.0000000000000000000000000000000000000000000003"),
		},
		{
			args: args{
				d:    mustNewFromString("0.000000000000000000000000000000000000000000901836202"),
				opts: New(WithRoundingPrecision(15), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("54703159147681578.1514852895273075959730711237955491690133829927977209580124"),
				opts: New(WithRoundingPrecision(59), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("54703159147681578.1514852895273075959730711237955491690133829927977209580124"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000000000000000000000005"),
				opts: New(WithRoundingPrecision(17), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000000000000000000000000000000039578053693375996216932325600263217353654"),
				opts: New(WithRoundingPrecision(42), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.000000000000000000000000000000039578053693"),
		},
		{
			args: args{
				d:    mustNewFromString("-400979013779505784551704647545324555644743917317817725"),
				opts: New(WithRoundingPrecision(51), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-400979013779505784551704647545324555644743917317817725"),
		},
		{
			args: args{
				d:    mustNewFromString("2549907485257905040787802731022172814.03247341030927871366393135398286301913246263649610180999011"),
				opts: New(WithRoundingPrecision(42), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("2549907485257905040787802731022172814.032473410309278713663931353982863019132463"),
		},
		{
			args: args{
				d:    mustNewFromString("7245563391265598645357861460253613932139592382610560614764364520097782949512752649"),
				opts: New(WithRoundingPrecision(40), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("7245563391265598645357861460253613932139592382610560614764364520097782949512752649"),
		},
		{
			args: args{
				d:    mustNewFromString("-64022435787355811014521281511793708435812347405139910972682589"),
				opts: New(WithRoundingPrecision(59), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-64022435787355811014521281511793708435812347405139910972682589"),
		},
		{
			args: args{
				d:    mustNewFromString("-76782672919180281245123823777032511965124724736456274885479622075418722"),
				opts: New(WithRoundingPrecision(48), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-76782672919180281245123823777032511965124724736456274885479622075418722"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000083135366316053183734543904737952651532784316140061929170739473518406297062533554026617147464"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("-273219046129778472266.0584854991093385965730749815531346353045049027336176088162559"),
				opts: New(WithRoundingPrecision(24), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-273219046129778472266.058485499109338596573075"),
		},
		{
			args: args{
				d:    mustNewFromString("6276464097096605785329824864148.52704981538099698591393138250952524233217779"),
				opts: New(WithRoundingPrecision(21), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("6276464097096605785329824864148.527049815380996985914"),
		},
		{
			args: args{
				d:    mustNewFromString("-597197.628834953506966767991553710700934413500204012426446876175175114500037146677042239668"),
				opts: New(WithRoundingPrecision(0), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-597198"),
		},
		{
			args: args{
				d:    mustNewFromString("-433359038877962603713455049783"),
				opts: New(WithRoundingPrecision(30), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-433359038877962603713455049783"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000000000000000000000000000000006381735336173415547900206847223271181528556195"),
				opts: New(WithRoundingPrecision(58), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.0000000000000000000000000000000006381735336173415547900207"),
		},
		{
			args: args{
				d:    mustNewFromString("-22678769.817248435493696742588538331241538369550386799148219117165563326051964281"),
				opts: New(WithRoundingPrecision(24), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-22678769.817248435493696742588538"),
		},
		{
			args: args{
				d:    mustNewFromString("5767307789536064608781837241295188919"),
				opts: New(WithRoundingPrecision(30), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("5767307789536064608781837241295188919"),
		},
		{
			args: args{
				d:    mustNewFromString("-88504154823150878334701258558002569539793415193610842759120001088201133334307983"),
				opts: New(WithRoundingPrecision(44), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-88504154823150878334701258558002569539793415193610842759120001088201133334307983"),
		},
		{
			args: args{
				d:    mustNewFromString("-329655464734888739743767364510089523323"),
				opts: New(WithRoundingPrecision(25), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-329655464734888739743767364510089523323"),
		},
		{
			args: args{
				d:    mustNewFromString("0.000000000000000000000845019203852002779189"),
				opts: New(WithRoundingPrecision(41), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0.00000000000000000000084501920385200277918"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00000000000000000000000000000000000000001462340018509"),
				opts: New(WithRoundingPrecision(21), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.000000000000000000001"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00000079078637152051768343295422323015639290504929"),
				opts: New(WithRoundingPrecision(42), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.000000790786371520517683432954223230156393"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0000000000054193378025297465767356639672580967150744942399"),
				opts: New(WithRoundingPrecision(50), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.00000000000541933780252974657673566396725809671507"),
		},
		{
			args: args{
				d:    mustNewFromString("-35620697798492911.066924841068786164325126832379766757683915930371102255700535220012496346147093317"),
				opts: New(WithRoundingPrecision(52), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-35620697798492911.0669248410687861643251268323797667576839159303711022"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000000000000000000005059713839821417238702105087169671933387005"),
				opts: New(WithRoundingPrecision(46), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.0000000000000000000050597138398214172387021051"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.81746270364999930607037613378763105641195817852303184573911882"),
				opts: New(WithRoundingPrecision(34), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-0.8174627036499993060703761337876311"),
		},
		{
			args: args{
				d:    mustNewFromString("26923162467831521466200388799932149017792464401239965995848900909703513553682"),
				opts: New(WithRoundingPrecision(59), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("26923162467831521466200388799932149017792464401239965995848900909703513553682"),
		},
		{
			args: args{
				d:    mustNewFromString("4554587644116353728395891927482"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("4554587644116353728395891927482"),
		},
		{
			args: args{
				d:    mustNewFromString("83"),
				opts: New(WithRoundingPrecision(1), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("83"),
		},
		{
			args: args{
				d:    mustNewFromString("-7097679626212584135194693334505819500.7627123978424311487730395375209597379059174819443305631091738"),
				opts: New(WithRoundingPrecision(41), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-7097679626212584135194693334505819500.76271239784243114877303953752095973790591"),
		},
		{
			args: args{
				d:    mustNewFromString("720941616590530465684319461159925340787620861616050215112729354513077297889437424470222725372.43418"),
				opts: New(WithRoundingPrecision(4), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("720941616590530465684319461159925340787620861616050215112729354513077297889437424470222725372.4341"),
		},
		{
			args: args{
				d:    mustNewFromString("7179182230007440380654240229988748528461622212340003478705"),
				opts: New(WithRoundingPrecision(16), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("7179182230007440380654240229988748528461622212340003478705"),
		},
		{
			args: args{
				d:    mustNewFromString("128138852434106311723518159896099183377408757231649238006509175039"),
				opts: New(WithRoundingPrecision(51), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("128138852434106311723518159896099183377408757231649238006509175039"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0000000000000000000000000034534834"),
				opts: New(WithRoundingPrecision(16), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.0000000000000001"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00000000000000000000082864928858923363665062475916780626021532507656936043414109352811732"),
				opts: New(WithRoundingPrecision(58), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.0000000000000000000008286492885892336366506247591678062602"),
		},
		{
			args: args{
				d:    mustNewFromString("1748715317929813133410156549170209422179478560908330825848622104018934659066"),
				opts: New(WithRoundingPrecision(57), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("1748715317929813133410156549170209422179478560908330825848622104018934659066"),
		},
		{
			args: args{
				d:    mustNewFromString("-554303811557294466.269761483473739646624314242607077815435340758612837177421989342652"),
				opts: New(WithRoundingPrecision(33), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-554303811557294466.269761483473739646624314242607077"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.19004485473016995992614957080209680408919713640428488619"),
				opts: New(WithRoundingPrecision(8), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-0.19004485"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000802214345771257141256247281416065552304304500535613033078792598113626175"),
				opts: New(WithRoundingPrecision(51), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.000080221434577125714125624728141606555230430450054"),
		},
		{
			args: args{
				d:    mustNewFromString("-6050582615205191601389958119203059837835097590785064613410822037914417495686026661"),
				opts: New(WithRoundingPrecision(31), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-6050582615205191601389958119203059837835097590785064613410822037914417495686026661"),
		},
		{
			args: args{
				d:    mustNewFromString("568254966593770.553753276551449605948238816764309803642928261672349658172008375162162314878680613"),
				opts: New(WithRoundingPrecision(21), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("568254966593770.553753276551449605949"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.00000000000000000000000089587927281631480798176250533957436898566513857011780218162097370714526"),
				opts: New(WithRoundingPrecision(44), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.00000000000000000000000089587927281631480799"),
		},
		{
			args: args{
				d:    mustNewFromString("-274912.8896024699118787839924993246206752520896053416203239133353705"),
				opts: New(WithRoundingPrecision(10), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-274912.8896024699"),
		},
		{
			args: args{
				d:    mustNewFromString("-2651358523359639"),
				opts: New(WithRoundingPrecision(16), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-2651358523359639"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000000000000000000000000000000000000112895858119340820153717620708673416"),
				opts: New(WithRoundingPrecision(45), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.000000000000000000000000000000000000112895859"),
		},
		{
			args: args{
				d:    mustNewFromString("842229243093860852173.0544396158817509837744408286148917213975696933283408713841831638764"),
				opts: New(WithRoundingPrecision(50), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("842229243093860852173.05443961588175098377444082861489172139756969332835"),
		},
		{
			args: args{
				d:    mustNewFromString("-699708233495.71227837422589896589188549642075052667001859282382939797996992686357419809583"),
				opts: New(WithRoundingPrecision(27), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-699708233495.712278374225898965891885496"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000000000716168920378590537721800581109521242491374877"),
				opts: New(WithRoundingPrecision(26), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.00000000071616892037859054"),
		},
		{
			args: args{
				d:    mustNewFromString("-2625562538887919963240549817430379735187837775384"),
				opts: New(WithRoundingPrecision(44), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-2625562538887919963240549817430379735187837775384"),
		},
		{
			args: args{
				d:    mustNewFromString("0.00000000000000000000002"),
				opts: New(WithRoundingPrecision(9), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("6067881766683695479556751950119377724336039886809300136812181462"),
				opts: New(WithRoundingPrecision(31), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("6067881766683695479556751950119377724336039886809300136812181462"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000000000000000000000000000000000000000772410227935606591033087412064412570098277"),
				opts: New(WithRoundingPrecision(51), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0.000000000000000000000000000000000000000077241022793"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0766240422992816781582074780208659294334413408481864862625859275536716954542357278357044523255"),
				opts: New(WithRoundingPrecision(50), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("0.07662404229928167815820747802086592943344134084819"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.000000000000000000000000000885255222637729340070545710310579917592457286140653"),
				opts: New(WithRoundingPrecision(33), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.000000000000000000000000000885256"),
		},
		{
			args: args{
				d:    mustNewFromString("55893020145100569857309952693924435456669213.356281068124371401302229274073839082240544043386519"),
				opts: New(WithRoundingPrecision(51), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("55893020145100569857309952693924435456669213.356281068124371401302229274073839082240544043386519"),
		},
		{
			args: args{
				d:    mustNewFromString("0.000000000044398802543439239843437872656117345711426288269229"),
				opts: New(WithRoundingPrecision(42), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0.000000000044398802543439239843437872656117"),
		},
		{
			args: args{
				d:    mustNewFromString("-51296363216658187515760473402291"),
				opts: New(WithRoundingPrecision(27), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-51296363216658187515760473402291"),
		},
		{
			args: args{
				d:    mustNewFromString("-272472379898107893040761485379027824396136886208"),
				opts: New(WithRoundingPrecision(40), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-272472379898107893040761485379027824396136886208"),
		},
		{
			args: args{
				d:    mustNewFromString("-19581413994383784948718328954653"),
				opts: New(WithRoundingPrecision(11), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-19581413994383784948718328954653"),
		},
		{
			args: args{
				d:    mustNewFromString("0.000000000000000000000025517761823978244291048210421988594612225022695964910425529"),
				opts: New(WithRoundingPrecision(21), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("0"),
		},
		{
			args: args{
				d:    mustNewFromString("663336.4573436793183219986595282312647796998714487327022132545955984591825466144183"),
				opts: New(WithRoundingPrecision(20), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("663336.45734367931832199866"),
		},
		{
			args: args{
				d:    mustNewFromString("-461221276.5204706386826154420723764419517018397461911607"),
				opts: New(WithRoundingPrecision(36), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("-461221276.52047063868261544207237644195170184"),
		},
		{
			args: args{
				d:    mustNewFromString("92664692270788697481952993240101"),
				opts: New(WithRoundingPrecision(9), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("92664692270788697481952993240101"),
		},
		{
			args: args{
				d:    mustNewFromString("0.0000000000000000417218487798321067688965201563233239322412080713783058725771499175637"),
				opts: New(WithRoundingPrecision(14), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.00000000000001"),
		},
		{
			args: args{
				d:    mustNewFromString("98445608185462908936594271820438358882244286228"),
				opts: New(WithRoundingPrecision(43), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("98445608185462908936594271820438358882244286228"),
		},
		{
			args: args{
				d:    mustNewFromString("0.000000000000000007838540274247171557849125417291807341828802631329086161031903106930206518314803525"),
				opts: New(WithRoundingPrecision(42), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("0.000000000000000007838540274247171557849126"),
		},
		{
			args: args{
				d:    mustNewFromString("-274462946120897177140732986024361620867165740004629583369799434.7526244978472107528071824755321251"),
				opts: New(WithRoundingPrecision(12), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("-274462946120897177140732986024361620867165740004629583369799434.752624497847"),
		},
		{
			args: args{
				d:    mustNewFromString("805407519521180265118391229"),
				opts: New(WithRoundingPrecision(27), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("805407519521180265118391229"),
		},
		{
			args: args{
				d:    mustNewFromString("4837.60752412303502513085517977565448616961234363618524491896285197928079211821305450283483280681236"),
				opts: New(WithRoundingPrecision(17), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("4837.60752412303502514"),
		},
		{
			args: args{
				d:    mustNewFromString("-0.0000000000000000002"),
				opts: New(WithRoundingPrecision(10), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("-0.0000000001"),
		},
		{
			args: args{
				d:    mustNewFromString("1758006742538130240388703498480688686072955030356.75241722336466855"),
				opts: New(WithRoundingPrecision(16), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("1758006742538130240388703498480688686072955030356.7524172233646685"),
		},
		{
			args: args{
				d:    mustNewFromString("71633752430127836728495483808.074842041489837644326088547938811892934755621628332271860178432369"),
				opts: New(WithRoundingPrecision(35), WithRoundingMode(constants.RoundHalfUp)),
			},
			want: mustNewFromString("71633752430127836728495483808.07484204148983764432608854793881189"),
		},
		{
			args: args{
				d:    mustNewFromString("5744453717566208360238616812981884481035389801"),
				opts: New(WithRoundingPrecision(46), WithRoundingMode(constants.RoundDown)),
			},
			want: mustNewFromString("5744453717566208360238616812981884481035389801"),
		},
		{
			args: args{
				d:    mustNewFromString("2988347090930431122495200201632971168964831173901728"),
				opts: New(WithRoundingPrecision(31), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("2988347090930431122495200201632971168964831173901728"),
		},
		{
			args: args{
				d:    mustNewFromString("2988347090930431122495200201632971168964831173901728"),
				opts: New(WithRoundingPrecision(31), WithRoundingMode(constants.RoundUp)),
			},
			want: mustNewFromString("2988347090930431122495200201632971168964831173901728"),
		},
	}
	for i, tt := range tests {
		got, err := DecimalRound(tt.args.d, tt.args.opts)
		if err != nil {
			t.Fatal(err)
		}

		if got.String() != tt.want.String() {
			t.Errorf("DecimalRound([%d]{d:[%+v], opts:[%+v]}) got = %v, want %v", i, tt.args.d.String(), tt.args.opts,
				got.String(), tt.want.String())
		}
	}
}
