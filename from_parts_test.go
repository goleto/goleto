package goleto

import (
	"testing"
	"time"
)

func TestNewBoleto(t *testing.T) {
	tests := []struct {
		name      string
		initArgs  []InitArg
		wantErr   bool
		wantValid string
	}{
		{
			name:      "zeroed boleto",
			wantValid: "00097000000000000000000000000000000000000000",
		},
		{
			name: "valid boleto - DST start date",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2024, time.January, 28, 0, 0, 0, 0, brTz), 2016, time.October, 16),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00194694900000010001234567890123456789012345",
		},
		{
			name: "valid boleto - before 2000-07-03",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2000, time.January, 28, 0, 0, 0, 0, brTz), 2000, time.July, 2),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00191099900000010001234567890123456789012345",
		},
		{
			name: "valid boleto",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2024, time.January, 28, 0, 0, 0, 0, brTz), 2023, time.December, 31),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00192958100000010001234567890123456789012345",
		},
		{
			name: "valid boleto 2",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2024, time.January, 28, 0, 0, 0, 0, brTz), 2025, time.February, 22),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00191100000000010001234567890123456789012345",
		},
		{
			name: "valid boleto in 1997",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2025, time.January, 28, 0, 0, 0, 0, brTz), 1997, time.October, 7),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00198000000000010001234567890123456789012345",
		},
		{
			name: "valid boleto in July 2 2000",
			initArgs: []InitArg{
				WithBankCode("001"),
				withExpirationDateAt(time.Date(2025, time.January, 28, 0, 0, 0, 0, brTz), 2000, time.July, 2),
				WithValue(1000),
				WithFreeField("1234567890123456789012345"),
			},
			wantErr:   false,
			wantValid: "00191099900000010001234567890123456789012345",
		},
		{
			name: "expiration date not representable",
			initArgs: []InitArg{
				withExpirationDateAt(time.Date(2025, time.January, 28, 0, 0, 0, 0, brTz), 2000, time.July, 3),
			},
			wantErr: true,
		},
		{
			name: "invalid bank code",
			initArgs: []InitArg{
				WithBankCode("1234"),
			},
			wantErr: true,
		},
		{
			name: "invalid expiration date",
			initArgs: []InitArg{
				WithExpirationDate(1800, time.January, 1),
			},
			wantErr: true,
		},
		{
			name: "value too large",
			initArgs: []InitArg{
				WithValue(10e10),
			},
			wantErr: true,
		},
		{
			name: "invalid free field",
			initArgs: []InitArg{
				WithFreeField("12345678901234567890123456"),
			},
			wantErr: true,
		},
		{
			name: "with random free field",
			initArgs: []InitArg{
				WithRandomFreeField(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boleto, err := NewBoleto(tt.initArgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoleto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if boleto.validBarcode != tt.wantValid {
					t.Errorf("NewBoleto() validBarcode = %v, want %v", boleto.validBarcode, tt.wantValid)
				}
				if _, err := ParseBoleto(boleto.validBarcode); err != nil {
					t.Error("NewBoleto() produced boleto has invalid check digit")
				}
			}
		})
	}
}

func TestWithRandomBankCode(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "valid random bank code",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			barcode := [44]byte{}
			err := WithRandomBankCode()(&barcode)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithRandomBankCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				bankCode := string(barcode[:3])
				if len(bankCode) != 3 || !digitsOnly(bankCode) {
					t.Errorf("WithRandomBankCode() produced invalid bank code: %v", bankCode)
				}
			}
		})
	}
}

func TestWithDateInAt(t *testing.T) {
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		now   time.Time
	}{
		{
			name:  "valid date range",
			start: time.Date(2023, time.January, 1, 0, 0, 0, 0, brTz),
			end:   time.Date(2023, time.December, 31, 0, 0, 0, 0, brTz),
			now:   time.Date(2023, time.July, 1, 0, 0, 0, 0, brTz),
		},
		{
			name:  "end before start",
			start: time.Date(2023, time.December, 31, 0, 0, 0, 0, brTz),
			end:   time.Date(2023, time.January, 1, 0, 0, 0, 0, brTz),
			now:   time.Date(2023, time.July, 1, 0, 0, 0, 0, brTz),
		},
		{
			name:  "same start and end date",
			start: time.Date(2023, time.July, 1, 0, 0, 0, 0, brTz),
			end:   time.Date(2023, time.July, 1, 0, 0, 0, 0, brTz),
			now:   time.Date(2023, time.July, 1, 0, 0, 0, 0, brTz),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boleto, err := NewBoleto(withDateInAt(tt.start, tt.end, tt.now))
			if err != nil {
				t.Errorf("withDateInAt() error = %v", err)
				return
			}

			year, month, day := boleto.calcExpirationDateAt(tt.now)

			expDate := time.Date(year, month, day, 0, 0, 0, 0, brTz)

			start, end := tt.start, tt.end
			if start.After(end) {
				start, end = end, start // Ensure start is before end
			}

			if expDate.Before(start) || expDate.After(end) {
				t.Errorf("withDateInAt() produced expiration date %v outside range [%v, %v]", expDate, tt.start, tt.end)
			}
		})
	}
}

func TestWithValueBy(t *testing.T) {
	tests := []struct {
		name    string
		mean    uint64
		stdDev  uint64
		wantErr bool
	}{
		{
			name:    "valid value within range",
			mean:    5000,
			stdDev:  1000,
			wantErr: false,
		},
		{
			name:    "mean close to upper limit",
			mean:    999999999,
			stdDev:  1000,
			wantErr: false,
		},
		{
			name:    "mean and stdDev resulting in out-of-range values",
			mean:    10e11,
			stdDev:  10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boleto, err := NewBoleto(WithValueBy(tt.mean, tt.stdDev))
			if (err != nil) != tt.wantErr {
				t.Errorf("WithValueBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				value := boleto.Value()
				if value > 10e8 {
					t.Errorf("WithValueBy() produced invalid value: %v", value)
				}
			}
		})
	}
}

func TestWithRandomFields(t *testing.T) {
	tests := []struct {
		name        string
		extraArgs   []InitArg
		extraAssert func(*testing.T, *Boleto)
		wantErr     bool
	}{
		{
			name:    "valid random fields",
			wantErr: false,
		},
		{
			name: "fixed value",
			extraArgs: []InitArg{
				WithValue(10000),
			},
			extraAssert: func(t *testing.T, boleto *Boleto) {
				if boleto.Value() != 10000 {
					t.Errorf("WithRandomFields() did not set fixed value correctly, got %v, want 10000", boleto.Value())
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []InitArg{WithRandomFields()}
			boleto, err := NewBoleto(append(args, tt.extraArgs...)...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithRandomFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.extraAssert != nil {
				tt.extraAssert(t, &boleto)
			}
		})
	}
}
