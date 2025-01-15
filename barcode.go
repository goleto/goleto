package goleto

func isValidBoletoBarcode(boleto string) bool {
	b := []byte(boleto)
	v := b[4] - 48
	return boletoDac11(b[0:4], b[5:]) == v
}

func isValidGdaBarcode(boleto string) bool {
	var checkFn func(b ...[]byte) uint8
	b := []byte(boleto)

	if b[0] != '8' {
		// Invalid product ID
		return false
	}

	if !(b[1] >= '1' && b[1] <= '9' && b[1] != '8') {
		// Invalid Segment ID
		return false
	}

	switch b[2] {
	case '6', '7':
		checkFn = dac10
	case '8', '9':
		checkFn = gdaDac11
	default:
		return false
	}

	return checkFn(b[0:3], b[4:44]) == b[3]-'0'
}
