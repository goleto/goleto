package goleto

import "errors"

// dac10 calculates the DAC10 (Modulo 10) check digit for the given chunks of bytes
// It takes one or more slices of bytes as input, where each byte represents a digit ('0'-'9').
// The function treats all chunks as a single concatenated string.
func dac10(chunks ...[]byte) uint8 {
	var sum int
	var pos int

	max := len(chunks) - 1
	for j := range chunks {
		b := chunks[max-j]

		max := len(b) - 1
		for i := range b {
			d := (2 - (pos & 1)) * int(b[max-i]-'0')
			if d > 9 {
				sum += d - 9
			} else {
				sum += d
			}
			pos++
		}
	}

	rem := uint8(sum % 10)

	if rem == 0 {
		return '0'
	} else {
		return 10 - rem + '0'
	}
}

// dac11 calculates a check digit using the DAC 11 algorithm.
// It takes one or more slices of bytes as input, where each byte represents a digit ('0'-'9').
// The function treats all chunks as a single concatenated string.
func dac11(adjustRem func(uint8) uint8, chunks ...[]byte) uint8 {
	var sum int
	var pos int

	max := len(chunks) - 1
	for j := range chunks {
		b := chunks[max-j]

		max := len(b) - 1
		for i := range b {
			sum += (2 + (pos & 7)) * int(b[max-i]-'0')
			pos++
		}
	}

	return adjustRem(uint8(sum%11)) + '0'
}

func boletoDac11(chunks ...[]byte) uint8 {
	return dac11(boletoDac11AdjustRem, chunks...)
}

func gdaDac11(chunks ...[]byte) uint8 {
	return dac11(gdaDac11AdjustRem, chunks...)
}

func boletoDac11AdjustRem(rem uint8) uint8 {
	switch rem {
	case 0, 1:
		return 1
	default:
		return 11 - rem
	}
}

func gdaDac11AdjustRem(rem uint8) uint8 {
	switch rem {
	case 0, 1:
		return 0
	case 10:
		return 1
	default:
		return 11 - rem
	}
}

func gdaCheckFn(b []byte) (checkFn func(b ...[]byte) uint8, err error) {
	// Infer the check digit algorithm from the third digit in barcode
	switch b[2] {
	case '6', '7': // when it's 6 or 7, it uses DAC 10
		checkFn = dac10
	case '8', '9': // when it's 6 or 7, it uses DAC 11
		checkFn = gdaDac11
	default: // other values are invalid
		checkFn = gdaDac11
		err = errors.New("invalid check fn")
	}
	return
}
