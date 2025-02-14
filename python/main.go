package main

import (
	_ "embed"
	"math"
	"math/rand"
	"strings"
	"time"
	"unsafe"

	"github.com/goleto/goleto"
)

// #include <string.h>
// #include <stdlib.h>
import "C"

//go:embed banks.tsv
var banks string
var bankCodes []string

//export RandomBoleto
func RandomBoleto(out *C.char) {
	for {
		b, err := randomBoleto()
		if err == nil {
			cstr := C.CString(b.Barcode())
			C.strcpy(out, cstr)
			C.free(unsafe.Pointer(cstr))
			return
		}
	}
}

func randomBoleto() (goleto.Boleto, error) {
	return goleto.NewBoleto(
		goleto.WithBankCode(randomBankCode()),
		goleto.WithExpirationDate(randomDate()),
		goleto.WithValue(randomValue()),
		goleto.WithFreeField(randomNumber(25)),
	)
}

func randomBankCode() string {
	return bankCodes[rand.Int31n(int32(len(bankCodes)))]
}

func randomDate() (int, time.Month, int) {
	var r float64
	for {
		r = rand.NormFloat64()*365 + 5
		if r < 0 {
			r = -rand.ExpFloat64() * 10
		}
		if math.Abs(r) < 4500 {
			break
		}
	}
	days := int(math.Round(r))

	return time.Now().AddDate(0, 0, int(days)).Date()
}

func randomValue() uint64 {
	var r float64
	for {
		r = rand.NormFloat64()*500 + 500
		if r > 0 && r < 10e8 {
			return uint64(math.RoundToEven(r * 100))
		}
	}
}

func randomNumber(n int) string {
	const digits = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Int31n(int32(len(digits)))]
	}
	return string(b)
}

func init() {
	for _, line := range strings.Split(banks, "\n") {
		parts := strings.Split(line, "\t")
		if len(parts) > 0 && len(parts[0]) == 3 {
			bankCodes = append(bankCodes, parts[0])
		}
	}
}

func main() {
}
