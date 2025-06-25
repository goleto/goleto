package goleto

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var ErrInvalidBankCode = errors.New("invalid bank code")
var ErrExpirationDateNotRepresentable = errors.New("expiration date not representable")
var ErrValueTooLarge = errors.New("value too large")
var ErrInvalidFreeField = errors.New("invalid free field")

type InitArg func(*[44]byte) error

func NewBoleto(initArgs ...InitArg) (boleto Boleto, err error) {
	barcode := [44]byte{
		'0', '0', '0', '9', '0', '0', '0', '0', '0', '0', '0',
		'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0',
		'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0',
		'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0',
	}

	var errs []error
	for _, init := range initArgs {
		if e := init(&barcode); e != nil {
			errs = append(errs, e)
		}
	}

	err = errors.Join(errs...)

	barcode[4] = boletoDac11(barcode[0:4], barcode[5:44])

	boleto.validBarcode = string(barcode[:])

	return
}

func WithBankCode(bankCode string) InitArg {
	return func(bp *[44]byte) error {
		if len(bankCode) > 3 || !digitsOnly(bankCode) {
			return ErrInvalidBankCode
		}

		_, _ = fmt.Fprintf(bytes.NewBuffer(bp[:0]), "%03s", bankCode)
		return nil
	}
}

func WithExpirationDate(year int, month time.Month, day int) InitArg {
	return withExpirationDateAt(time.Now(), year, month, day)
}

func withExpirationDateAt(now time.Time, year int, month time.Month, day int) InitArg {
	return func(bp *[44]byte) error {
		epoch := time.Date(2000, time.July, 3, 0, 0, 0, 0, brTz)

		expirationDate := time.Date(year, month, day, 0, 0, 0, 0, brTz)

		if expirationDate.Before(epoch) {
			expirationDate = expirationDate.Add(-12 * time.Hour)
		} else {
			expirationDate = expirationDate.Add(12 * time.Hour)
		}

		d := expirationDate.Sub(epoch)
		day := 24 * time.Hour
		factor := int64(d / day)

		if factor < -1000 {
			return ErrExpirationDateNotRepresentable
		}

		if factor >= 0 {
			today := now.In(brTz)
			if today.IsDST() {
				today = today.Add(time.Hour)
			}
			daysSinceEpoch := int64(today.Sub(epoch) / (24 * time.Hour))

			diff := factor - daysSinceEpoch

			if diff > 4500 || diff <= -4500 {
				return ErrExpirationDateNotRepresentable
			}

			factor %= 9000
		}

		_, _ = fmt.Fprintf(bytes.NewBuffer(bp[:5]), "%04d", factor+1000)

		return nil
	}
}

func WithValue(value uint64) InitArg {
	return func(bp *[44]byte) error {
		if value >= 10e10 {
			return ErrValueTooLarge
		}

		_, _ = fmt.Fprintf(bytes.NewBuffer(bp[:9]), "%0*d", 10, value)

		return nil
	}
}

func WithFreeField(freeField string) InitArg {
	return func(bp *[44]byte) error {
		if len(freeField) > 25 || !digitsOnly(freeField) {
			return ErrInvalidFreeField
		}

		_, _ = fmt.Fprintf(bytes.NewBuffer(bp[:19]), "%0*s", 25, freeField)

		return nil
	}
}

func WithRandomFields() InitArg {
	inits := []InitArg{
		WithRandomBankCode(),
		WithDateIn(time.Now().AddDate(0, 0, 1), time.Now().AddDate(0, 0, 30)),
		WithValueBy(3000, 1000),
		WithRandomFreeField(),
	}

	return func(bp *[44]byte) error {
		for _, init := range inits {
			if e := init(bp); e != nil {
				return e
			}
		}

		return nil
	}
}

func WithRandomBankCode() InitArg {
	return func(bp *[44]byte) error {
		bankCode := BankingCodes[rand.Int31n(int32(len(BankingCodes)))]
		return WithBankCode(bankCode)(bp)
	}
}

func WithDateIn(start, end time.Time) InitArg {
	return withDateInAt(start, end, time.Now())
}

func withDateInAt(start, end, now time.Time) InitArg {
	if end.Before(start) {
		start, end = end, start // Ensure start is before end
	}
	return func(bp *[44]byte) error {
		var randomInstant int64

		if r := end.Unix() - start.Unix(); r == 0 {
			randomInstant = start.Unix()
		} else {
			randomInstant = start.Unix() + rand.Int63n(r)
		}

		year, month, day := time.Unix(randomInstant, 0).Date()
		return withExpirationDateAt(now, year, month, day)(bp)
	}
}

func WithValueBy(mean uint64, stdDev uint64) InitArg {
	return func(bp *[44]byte) error {
		var r float64
		for range 100 {
			r = rand.NormFloat64()*float64(stdDev) + float64(mean)
			if r > 0 && r < 10e8 {
				return WithValue(uint64(math.RoundToEven(r)))(bp)
			}
		}

		return ErrValueTooLarge
	}
}

func WithRandomFreeField() InitArg {
	return func(bp *[44]byte) error {
		const digits = "0123456789"
		b := make([]byte, 25)
		for i := range b {
			b[i] = digits[rand.Int31n(int32(len(digits)))]
		}
		return WithFreeField(string(b))(bp)
	}
}
