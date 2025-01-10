package goleto

import (
	"unsafe"
)

func isValidBarcode(boleto string) bool {
	b := *(*[]byte)(unsafe.Pointer(&boleto))
	v := b[4] - 48
	return calcBarcodeCheckDigit(boleto) == v
}

func calcBarcodeCheckDigit(boleto string) uint8 {
	var sum int
	b := *(*[]byte)(unsafe.Pointer(&boleto))

	sum += 4 * int(b[0]-'0')
	sum += 3 * int(b[1]-'0')
	sum += 2 * int(b[2]-'0')
	sum += 9 * int(b[3]-'0')

	sum += 8 * int(b[5]-'0')
	sum += 7 * int(b[6]-'0')
	sum += 6 * int(b[7]-'0')
	sum += 5 * int(b[8]-'0')
	sum += 4 * int(b[9]-'0')
	sum += 3 * int(b[10]-'0')
	sum += 2 * int(b[11]-'0')

	sum += 9 * int(b[12]-'0')
	sum += 8 * int(b[13]-'0')
	sum += 7 * int(b[14]-'0')
	sum += 6 * int(b[15]-'0')
	sum += 5 * int(b[16]-'0')
	sum += 4 * int(b[17]-'0')
	sum += 3 * int(b[18]-'0')
	sum += 2 * int(b[19]-'0')

	sum += 9 * int(b[20]-'0')
	sum += 8 * int(b[21]-'0')
	sum += 7 * int(b[22]-'0')
	sum += 6 * int(b[23]-'0')
	sum += 5 * int(b[24]-'0')
	sum += 4 * int(b[25]-'0')
	sum += 3 * int(b[26]-'0')
	sum += 2 * int(b[27]-'0')

	sum += 9 * int(b[28]-'0')
	sum += 8 * int(b[29]-'0')
	sum += 7 * int(b[30]-'0')
	sum += 6 * int(b[31]-'0')
	sum += 5 * int(b[32]-'0')
	sum += 4 * int(b[33]-'0')
	sum += 3 * int(b[34]-'0')
	sum += 2 * int(b[35]-'0')

	sum += 9 * int(b[36]-'0')
	sum += 8 * int(b[37]-'0')
	sum += 7 * int(b[38]-'0')
	sum += 6 * int(b[39]-'0')
	sum += 5 * int(b[40]-'0')
	sum += 4 * int(b[41]-'0')
	sum += 3 * int(b[42]-'0')
	sum += 2 * int(b[43]-'0')

	mod := uint8(11 - (sum % 11))

	switch mod {
	case 0, 10, 11:
		return 1
	default:
		return mod
	}
}
