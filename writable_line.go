package goleto

func (Boleto) writableLineLength() int {
	return 47
}

func (Gda) writableLineLength() int {
	return 48
}

func (Boleto) isValidWritableLine(writableLine string) bool {
	b := []byte(writableLine)

	return dac10(b[0:9]) == b[9] &&
		dac10(b[10:20]) == b[20] &&
		dac10(b[21:31]) == b[31]
}

func (Gda) isValidWritableLine(writableLine string) bool {
	b := []byte(writableLine)

	if checkFn, err := gdaCheckFn(b); err == nil {
		return checkFn(b[0:11]) == b[11] &&
			checkFn(b[12:23]) == b[23] &&
			checkFn(b[24:35]) == b[35] &&
			checkFn(b[36:47]) == b[47]
	}

	return false
}
