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

	sum += 4 * int(b[0]-48)
	sum += 3 * int(b[1]-48)
	sum += 2 * int(b[2]-48)
	sum += 9 * int(b[3]-48)

	sum += 8 * int(b[5]-48)
	sum += 7 * int(b[6]-48)
	sum += 6 * int(b[7]-48)
	sum += 5 * int(b[8]-48)
	sum += 4 * int(b[9]-48)
	sum += 3 * int(b[10]-48)
	sum += 2 * int(b[11]-48)

	sum += 9 * int(b[12]-48)
	sum += 8 * int(b[13]-48)
	sum += 7 * int(b[14]-48)
	sum += 6 * int(b[15]-48)
	sum += 5 * int(b[16]-48)
	sum += 4 * int(b[17]-48)
	sum += 3 * int(b[18]-48)
	sum += 2 * int(b[19]-48)

	sum += 9 * int(b[20]-48)
	sum += 8 * int(b[21]-48)
	sum += 7 * int(b[22]-48)
	sum += 6 * int(b[23]-48)
	sum += 5 * int(b[24]-48)
	sum += 4 * int(b[25]-48)
	sum += 3 * int(b[26]-48)
	sum += 2 * int(b[27]-48)

	sum += 9 * int(b[28]-48)
	sum += 8 * int(b[29]-48)
	sum += 7 * int(b[30]-48)
	sum += 6 * int(b[31]-48)
	sum += 5 * int(b[32]-48)
	sum += 4 * int(b[33]-48)
	sum += 3 * int(b[34]-48)
	sum += 2 * int(b[35]-48)

	sum += 9 * int(b[36]-48)
	sum += 8 * int(b[37]-48)
	sum += 7 * int(b[38]-48)
	sum += 6 * int(b[39]-48)
	sum += 5 * int(b[40]-48)
	sum += 4 * int(b[41]-48)
	sum += 3 * int(b[42]-48)
	sum += 2 * int(b[43]-48)

	mod := uint8(11 - (sum % 11))

	switch mod {
	case 0, 10, 11:
		return 1
	default:
		return mod
	}
}
