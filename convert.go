package goleto

import (
	"strings"
	"unsafe"
)

// writableLineToBarcode converts a writable line string to a barcode string.
func writableLineToBarcode(writableLine string) string {
	var b strings.Builder

	b.Grow(44)

	b.WriteString(writableLine[0:4])   // bank // currency
	b.WriteByte(writableLine[32])      // check digit
	b.WriteString(writableLine[33:47]) // fator de vencimento // valor

	// free field
	b.WriteString(writableLine[4:9])
	b.WriteString(writableLine[10:20])
	b.WriteString(writableLine[21:31])

	return b.String()
}

// barcodeToWritableLine converts a barcode string to a writable line format.
func barcodeToWritableLine(barcode string) string {
	var sb strings.Builder

	sb.Grow(47)

	sb.WriteString(barcode[:4])
	sb.WriteString(barcode[19:24])
	sb.WriteByte('0') // first check digit

	sb.WriteString(barcode[24:34])
	sb.WriteByte('0') // second check digit

	sb.WriteString(barcode[34:44])
	sb.WriteByte('0') // third check digit

	sb.WriteString(barcode[4:19])

	s := sb.String()
	b := *(*[]byte)(unsafe.Pointer(&s))

	b[9] = calcWritableLineFieldCheckDigit(b[0:9]) + '0'
	b[20] = calcWritableLineFieldCheckDigit(b[10:20]) + '0'
	b[31] = calcWritableLineFieldCheckDigit(b[21:31]) + '0'

	return s
}
