package number

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func New(opt ...Option) *Options {
	opts := *defaultOptions

	for _, o := range opt {
		o.apply(&opts)
	}

	return &opts
}

// DecimalFormat produces a string form of the given decimal.Decimal in base 10
//
// ref: https://github.com/dustin/go-humanize/blob/master/commaf.go#L13
func DecimalFormat(d decimal.Decimal, opts *Options) (formatted string) {
	defer func() {
		formatted = removeNonBreakingSpace(formatted)
	}()

	buf := &bytes.Buffer{}
	s := d.String()
	parts := strings.Split(s, ".")
	if opts.decimalPlaces != nil && len(parts) > 1 && *opts.decimalPlaces < uint(len(parts[1])) {
		s = d.StringFixed(int32(*opts.decimalPlaces))
	}

	if d.Sign() < 0 {
		s = s[1:]
		buf.WriteByte('-')
		d.Abs()
	}

	parts = strings.Split(s, ".")
	buf.WriteString(formatGroup(parts[0], opts.groupSize, opts.secondaryGroupSize, opts.groupSeparator))
	if len(parts) == 1 && (opts.decimalPlaces == nil || *opts.decimalPlaces == 0) {
		return buf.String()
	}

	if opts.decimalPlaces != nil {
		if len(parts) == 1 {
			parts = append(parts, "")
		}

		if size := *opts.decimalPlaces - uint(len(parts[1])); size > 0 {
			parts[1] += fmt.Sprintf("%0*d", size, 0)
		}

		parts[1] = parts[1][:*opts.decimalPlaces]
	}

	buf.WriteByte(opts.decimalSeparator)
	buf.WriteString(formatFraction(parts[1], opts.fractionGroupSize, opts.fractionGroupSeparator))

	return strings.TrimRight(buf.String(), string(opts.fractionGroupSeparator))
}

func formatGroup(num string, groupSize, secondaryGroupSize uint, groupSeparator byte) string {
	var buf = new(bytes.Buffer)
	var pos uint = 0
	iLen := uint(len(num))
	if groupSize > 1 {
		if groupSize < iLen {
			iLen -= groupSize
			if secondaryGroupSize > 0 {
				groupSize = secondaryGroupSize
			}

			if subPOS := iLen % groupSize; subPOS != 0 {
				pos += subPOS
				buf.WriteString(num[:pos])
				buf.WriteByte(groupSeparator)
			}

			for ; pos < iLen; pos += groupSize {
				buf.WriteString(num[pos : pos+groupSize])
				buf.WriteByte(groupSeparator)
			}

			buf.WriteString(num[iLen:])
		} else {
			buf.WriteString(num)
		}
	}

	return buf.String()
}

func formatFraction(num string, fractionGroupSize uint, fractionGroupSeparator byte) string {
	var buf = new(bytes.Buffer)
	var pos uint = 0
	fLen := uint(len(num))
	if fractionGroupSize == 0 {
		buf.WriteString(num)
		return buf.String()
	}

	lastPOS := fLen % fractionGroupSize
	for ; pos < fLen-lastPOS; pos += fractionGroupSize {
		buf.WriteString(num[pos : pos+fractionGroupSize])
		buf.WriteByte(fractionGroupSeparator)
	}

	if lastPOS > 0 {
		buf.WriteString(num[pos : pos+lastPOS])
		buf.WriteByte(fractionGroupSeparator)
	}

	return buf.String()
}

// DecimalRound sets d to its value rounded to the given precision using the given rounding mode.
//
// Returns d, which was modified in place.
func DecimalRound(d decimal.Decimal, opts *Options) (decimal.Decimal, error) {
	if !opts.mode.Valid() {
		return decimal.Decimal{}, ErrInvalidRM
	}

	return modeHandles[opts.mode](d, opts.prec)
}
