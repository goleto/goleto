package goleto

func (Boleto) isValidBarcode(barcode string) bool {
	b := []byte(barcode)
	return boletoDac11(b[0:4], b[5:44]) == b[4]
}

func (Gda) isValidBarcode(barcode string) bool {
	b := []byte(barcode)

	// The first digit in barcode is the product ID, 8 is the only valid value
	productId := b[0]
	if productId != '8' {
		// Invalid product ID
		return false
	}

	// The second digit in barcode is the segment ID, 0 and 8 are invalid
	segmentId := b[1]
	if !(segmentId >= '1' && segmentId <= '9' && segmentId != '8') {
		// Invalid Segment ID
		return false
	}

	checkFn, err := gdaCheckFn(b)
	if err != nil {
		return false
	}

	return checkFn(b[0:3], b[4:44]) == b[3]
}

func (b *Boleto) setValidBarcode(validBarcode string) {
	b.validBarcode = validBarcode
}

func (g *Gda) setValidBarcode(validBarcode string) {
	g.validBarcode = validBarcode
}
