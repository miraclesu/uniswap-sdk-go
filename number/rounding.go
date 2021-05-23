package number

import (
	"errors"

	"github.com/shopspring/decimal"
	gorounding "github.com/wadey/go-rounding"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

type modeHandler func(decimal.Decimal, int) (decimal.Decimal, error)

var (
	modeHandles = map[constants.Rounding]modeHandler{
		constants.RoundDown:   roundDownHandle,
		constants.RoundHalfUp: roundHalfUpHandle,
		constants.RoundUp:     roundUpHandle,
	}

	// ErrInvalidRM invalid rounding mode
	ErrInvalidRM = errors.New("invalid rounding mode")
)

func roundDownHandle(d decimal.Decimal, prec int) (decimal.Decimal, error) {
	return round(d, prec, gorounding.Down)
}

func roundHalfUpHandle(d decimal.Decimal, prec int) (decimal.Decimal, error) {
	return round(d, prec, gorounding.HalfUp)
}

func roundUpHandle(d decimal.Decimal, prec int) (decimal.Decimal, error) {
	return round(d, prec, gorounding.Up)
}

func round(d decimal.Decimal, prec int, mode gorounding.RoundingMode) (decimal.Decimal, error) {
	return decimal.NewFromString(gorounding.Round(d.Rat(), prec, mode).FloatString(prec))
}
