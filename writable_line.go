package goleto

func isValidBoletoWritableLine(writableLine string) bool {
	b := []byte(writableLine)

	return dac10(b[0:9]) == b[9]-'0' &&
		dac10(b[10:20]) == b[20]-'0' &&
		dac10(b[21:31]) == b[31]-'0'
}

func isValidGdaWritableLine(writableLine string) bool {
	var checkFn func(b ...[]byte) uint8
	b := []byte(writableLine)

	switch b[2] {
	case '6', '7':
		checkFn = dac10
	case '8', '9':
		checkFn = gdaDac11
	default:
		return false
	}

	return checkFn(b[0:11]) == b[11]-'0' &&
		checkFn(b[12:23]) == b[23]-'0' &&
		checkFn(b[24:35]) == b[35]-'0' &&
		checkFn(b[36:47]) == b[47]-'0'
}
