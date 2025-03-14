package goleto

import (
	"fmt"
	"strconv"
	"time"
)

type Boleto struct {
	validBarcode string
}

// Barcode returns the Boleto on bar code format.
func (b Boleto) Barcode() string {
	return b.validBarcode
}

// WritableLine returns the Boleto on writable line format.
func (b Boleto) WritableLine() string {
	return b.barcodeToWritableLine(b.validBarcode)
}

// BankCode returns the bank code extracted from the valid barcode of the Boleto.
//
// The bank code is represented by the first three characters of the barcode.
func (b Boleto) BankCode() string {
	return b.validBarcode[0:3]
}

// CurrencyCode returns the currency code of the Boleto as a string.
//
// The currency code is extracted from the validBarcode field of the Boleto,
// specifically from the 4th character (index 3).
func (b Boleto) CurrencyCode() string {
	return b.validBarcode[3:4]
}

// DateFactor extracts the date factor from the Boleto.
//
// The date factor is a 4-digit substring starting at the 6th character of the barcode.
// It converts this substring to an unsigned 16-bit integer and returns it.
func (b Boleto) DateFactor() uint16 {
	factor, _ := strconv.ParseUint(b.validBarcode[5:9], 10, 16)
	return uint16(factor)
}

// ExpirationDate calculates and returns the expiration date of the Boleto.
//
// It extracts the number of days from the validBarcode field, adds these days
// to the base date of August 7, 1997, and returns the resulting year, month, and day.
//
// Returns:
//
//	year  - the year of the expiration date
//	month - the month of the expiration date
//	day   - the day of the expiration date
func (b Boleto) ExpirationDate() (year int, month time.Month, day int) {
	return b.calcExpirationDateAt(time.Now())
}

var brTz *time.Location

func (b Boleto) calcExpirationDateAt(now time.Time) (year int, month time.Month, day int) {
	factor := int64(b.DateFactor())

	if factor < 1000 {
		epoch := time.Date(1997, time.October, 7, 0, 0, 0, 0, brTz)
		return epoch.AddDate(0, 0, int(factor)).Date()
	}

	epoch := time.Date(2000, time.July, 3, 0, 0, 0, 0, brTz)

	today := now.In(brTz)
	if today.IsDST() {
		today = today.Add(time.Hour)
	}
	daysSinceEpoch := int64(today.Sub(epoch) / (24 * time.Hour))
	epochAdjust := (daysSinceEpoch % 9000) - (factor - 1000)

	if daysSinceEpoch >= 4500 {
		if epochAdjust >= 4500 {
			epochAdjust -= 9000
		} else if epochAdjust < -4500 {
			epochAdjust += 9000
		}
	}

	// The date can be the start of a DST period. So, e.g.:
	//
	// 1997-08-07T00:00:00 + 6949 days = 2016-10-16T00:00:00.
	//
	// But at that date, a DLS period starts. So the go time library go converts
	// it to 2016-10-15H23:00:00. To avoid that day shift, we add 12 hours to
	// the final calculated date.
	return epoch.AddDate(0, 0, int(daysSinceEpoch-epochAdjust)).Add(12 * time.Hour).Date()
}

// Value extracts and returns the monetary value of the Boleto in cents.
//
// It parses the value from the validBarcode field, which is expected to be
// a string of digits. The value is extracted from the 10th to the 19th character
// of the validBarcode string and converted to an unsigned 64-bit integer.
func (b Boleto) Value() uint64 {
	cents, _ := strconv.ParseUint(b.validBarcode[9:19], 10, 64)
	return cents
}

// FreeField extracts and returns a specific portion of the valid barcode
// from the Boleto struct as a string. The portion extracted is from the 20th to
// 44th  character of the validBarcode field. This data is opaque and handled by each bank.
func (b Boleto) FreeField() string {
	return b.validBarcode[19:44]
}

func init() {
	var err error
	brTz, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(fmt.Sprintf("cannot load Brasília time zone: %v", err))
	}
}
