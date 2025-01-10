package goleto

import (
	"unsafe"
)

func isValidWritableLine(writableLine string) bool {
	b := *(*[]byte)(unsafe.Pointer(&writableLine))

	return calcWritableLineFieldCheckDigit(b[0:9]) == b[9]-'0' &&
		calcWritableLineFieldCheckDigit(b[10:20]) == b[20]-'0' &&
		calcWritableLineFieldCheckDigit(b[21:31]) == b[31]-'0'
}

func calcWritableLineFieldCheckDigit(b []byte) uint8 {
	var sum int
	max := len(b) - 1
	for i := range b {
		d := (2 - (i & 1)) * int(b[max-i]-'0')
		if d > 9 {
			sum += d - 9
		} else {
			sum += d
		}
	}

	mod := uint8(sum % 10)

	if mod == 0 {
		return 0
	} else {
		return 10 - mod
	}
}
