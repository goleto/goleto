package goleto

import (
	"errors"
	"unicode"
)

// ErrInvalidCode is returned when the provided barcode or writable line does not meet the expected format or criteria.
var ErrInvalidCode = errors.New("invalid code")

// ParseBoleto takes a string `s` representing a boleto (Brazilian payment slip) and
// returns a Boleto struct if the string is a valid barcode or writable line.
//
// Parameters:
//   - s: A barcode or writable line string.
//
// Returns:
//   - Boleto: A struct representing the valid boleto.
//   - error: An error if `s` is not a valid barcode or writable line, specifically ErrInvalidCode.
func ParseBoleto(s string) (Boleto, error) {
	return parse[Boleto](s)
}

// ParseGda parses a given string to create a Gda (Guia de Arrecadação) if the string is a valid GDA barcode or writable line.
// The function checks the length and the format of the string to determine its validity.
// It returns a Gda object and nil error if the string is valid, otherwise it returns an empty Gda object and an ErrInvalidBarcode error.
//
// Parameters:
//   - s: A barcode or writable line string.
//
// Returns:
//   - Gda: The parsed Gda object if the string is valid.
//   - error: An error indicating the string is not a valid GDA, specifically ErrInvalidBarcode.
func ParseGda(s string) (Gda, error) {
	return parse[Gda](s)
}

type parsable interface {
	writableLineLength() int
	isValidWritableLine(string) bool
	writableLineToBarcode(string) string
	isValidBarcode(string) bool
}

type updatable[T any] interface {
	setValidBarcode(string)
	*T
}

func parse[P parsable, PP updatable[P]](s string) (p P, err error) {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			err = ErrInvalidCode
			return
		}
	}

	switch len(s) {
	case p.writableLineLength():
		if !p.isValidWritableLine(s) {
			break
		}
		s = p.writableLineToBarcode(s)
		fallthrough
	case 44:
		if p.isValidBarcode(s) {
			PP(&p).setValidBarcode(s)
			return
		}
	}
	err = ErrInvalidCode
	return
}
