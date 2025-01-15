package goleto

import (
	"errors"
	"unicode"
)

// ErrInvalidBarcode is returned when the provided barcode does not meet the expected format or criteria.
var ErrInvalidBarcode = errors.New("invalid barcode")

// ParseBoleto takes a string `s` representing a boleto (Brazilian payment slip) and
// returns a Boleto struct if the string is a valid barcode or writable line.
//
// Parameters:
//   - s: A barcode or writable line string.
//
// Returns:
//   - Boleto: A struct representing the valid boleto.
//   - error: An error if the boleto is invalid, specifically ErrInvalidBarcode.
func ParseBoleto(s string) (Boleto, error) {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return Boleto{}, ErrInvalidBarcode
		}
	}

	switch len(s) {
	case 47:
		if !isValidBoletoWritableLine(s) {
			break
		}
		s = writableLineToBoletoBarcode(s)
		fallthrough
	case 44:
		if isValidBoletoBarcode(s) {
			return Boleto{s}, nil
		}
	}
	return Boleto{}, ErrInvalidBarcode
}

// ParseGda parses a given string to create a Gda (Guia de Arrecadação) if the string is a valid GDA barcode or writable line.
// The function checks the length and the format of the string to determine its validity.
// It returns a Gda object and nil error if the string is valid, otherwise it returns an empty Gda object and an ErrInvalidBarcode error.
//
// Parameters:
//   - s: The string to be parsed.
//
// Returns:
//   - Gda: The parsed Gda object if the string is valid.
//   - error: An error indicating the string is not a valid GDA, specifically ErrInvalidBarcode.
func ParseGda(s string) (Gda, error) {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return Gda{}, ErrInvalidBarcode
		}
	}

	switch len(s) {
	case 48:
		if !isValidGdaWritableLine(s) {
			break
		}
		s = writableLineToGdaBarcode(s)
		fallthrough
	case 44:
		if isValidGdaBarcode(s) {
			return Gda{s}, nil
		}
	}
	return Gda{}, ErrInvalidBarcode
}
