package goleto

import (
	"errors"
	"unicode"
)

// ErrInvalidBarcode is returned when the provided barcode does not meet the expected format or criteria.
var ErrInvalidBarcode = errors.New("invalid barcode")

// Parse takes a string `s` representing a boleto (Brazilian payment slip) and
// returns a Boleto struct if the string is a valid barcode or writable line.
//
// Parameters:
//   - s: A barcode or writable line string.
//
// Returns:
//   - Boleto: A struct representing the valid boleto.
//   - error: An error if the boleto is invalid, specifically ErrInvalidBarcode.
func Parse(s string) (Boleto, error) {

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return Boleto{}, ErrInvalidBarcode
		}
	}

	switch len(s) {
	case 47:
		if !isValidWritableLine(s) {
			break
		}
		s = writableLineToBarcode(s)
		fallthrough
	case 44:
		if isValidBarcode(s) {
			return Boleto{s}, nil
		}
	}
	return Boleto{}, ErrInvalidBarcode
}
