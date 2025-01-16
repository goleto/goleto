package goleto

import (
	"strings"
)

func (Boleto) writableLineToBarcode(writableLine string) string {
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

func (Gda) writableLineToBarcode(writableLine string) string {
	var b strings.Builder

	b.Grow(44)

	b.WriteString(writableLine[0:11])
	b.WriteString(writableLine[12:23])
	b.WriteString(writableLine[24:35])
	b.WriteString(writableLine[36:47])

	return b.String()
}

func (Boleto) barcodeToWritableLine(barcode string) string {
	var b [47]byte

	copy(b[0:4], barcode[0:4])
	copy(b[4:9], barcode[19:24])

	copy(b[10:20], barcode[24:34])
	copy(b[21:31], barcode[34:44])

	copy(b[32:47], barcode[4:19])

	b[9] = dac10(b[0:9])
	b[20] = dac10(b[10:20])
	b[31] = dac10(b[21:31])

	return string(b[:])
}

func (Gda) barcodeToWritableLine(barcode string) string {
	var b [48]byte

	copy(b[0:11], barcode[0:11])
	copy(b[12:23], barcode[11:22])
	copy(b[24:35], barcode[22:33])
	copy(b[36:47], barcode[33:44])

	checkFn, _ := gdaCheckFn(b[:])
	b[11] = checkFn(b[0:11])
	b[23] = checkFn(b[12:23])
	b[35] = checkFn(b[24:35])
	b[47] = checkFn(b[36:47])

	return string(b[:])
}
