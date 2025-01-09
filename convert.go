package goleto

import "strings"

// writableLineToBarcode converts a writable line string to a barcode string.
func writableLineToBarcode(writableLine string) string {
	b := strings.Builder{}

	b.WriteString(writableLine[0:4])   // bank // currency
	b.WriteByte(writableLine[32])      // check digit
	b.WriteString(writableLine[33:47]) // fator de vencimento // valor

	// free field
	b.WriteString(writableLine[4:9])
	b.WriteString(writableLine[10:20])
	b.WriteString(writableLine[21:31])

	return b.String()
}

// barcodeToWritableLine converts a barcode string into a writable line format.
func barcodeToWritableLine(_barcode string) string {
	// TODO
	return ""
}
