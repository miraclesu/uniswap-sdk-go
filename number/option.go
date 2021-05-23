package number

import (
	"bytes"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

type (
	Options struct {
		formatOptions
		roundingOptions
	}

	formatOptions struct {
		decimalSeparator       byte
		groupSeparator         byte
		groupSize              uint
		secondaryGroupSize     uint
		fractionGroupSeparator byte
		fractionGroupSize      uint
		decimalPlaces          *uint
	}

	roundingOptions struct {
		mode constants.Rounding
		prec int
	}
)

var (
	defaultOptions = &Options{
		formatOptions:   defaultFormatOptions,
		roundingOptions: defaultRoundingOptions,
	}

	defaultFormatOptions = formatOptions{
		decimalSeparator:       '.',
		groupSeparator:         ',',
		groupSize:              3,
		secondaryGroupSize:     0,
		fractionGroupSeparator: '\xA0',
		fractionGroupSize:      0,
		decimalPlaces:          nil,
	}

	defaultRoundingOptions = roundingOptions{
		mode: constants.RoundHalfUp,
		prec: -1,
	}
)

type Option interface {
	apply(*Options)
}

type funcOption struct {
	f func(*Options)
}

func (fo *funcOption) apply(do *Options) {
	fo.f(do)
}

func (o *Options) Apply(opt ...Option) {
	for _, of := range opt {
		of.apply(o)
	}
}

func newFuncOption(f func(*Options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithDecimalSeparator(decimalSeparator byte) Option {
	return newFuncOption(func(o *Options) {
		o.decimalSeparator = decimalSeparator
	})
}

func WithGroupSeparator(groupSeparator byte) Option {
	return newFuncOption(func(o *Options) {
		o.groupSeparator = groupSeparator
	})
}

func WithGroupSize(groupSize uint) Option {
	return newFuncOption(func(o *Options) {
		o.groupSize = groupSize
	})
}

func WithSecondaryGroupSize(secondaryGroupSize uint) Option {
	return newFuncOption(func(o *Options) {
		o.secondaryGroupSize = secondaryGroupSize
	})
}

func WithFractionGroupSeparator(fractionGroupSeparator byte) Option {
	return newFuncOption(func(o *Options) {
		o.fractionGroupSeparator = fractionGroupSeparator
	})
}

func WithFractionGroupSize(fractionGroupSize uint) Option {
	return newFuncOption(func(o *Options) {
		o.fractionGroupSize = fractionGroupSize
	})
}

func WithDecimalPlaces(decimalPlaces uint) Option {
	return newFuncOption(func(o *Options) {
		o.decimalPlaces = &decimalPlaces
	})
}

func WithRoundingMode(mode constants.Rounding) Option {
	return newFuncOption(func(o *Options) {
		o.mode = mode
	})
}

func WithRoundingPrecision(precision int) Option {
	return newFuncOption(func(o *Options) {
		o.prec = precision
	})
}

// remove byte '\xA0'
func removeNonBreakingSpace(s string) string {
	return string(bytes.ReplaceAll([]byte(s), []byte{'\xA0'}, []byte("")))
}
