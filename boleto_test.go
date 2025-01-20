package goleto

import (
	"testing"
	"time"
)

func TestBoleto(t *testing.T) {
	tests := []struct {
		name                 string
		barcode              string
		expectedBank         string
		expectedCurr         string
		expectedYear         int
		expectedMonth        time.Month
		expectedDay          int
		expectedValue        uint64
		expectedFree         string
		expectedWritableLine string
	}{
		{
			name:                 "Valid Boleto 1",
			barcode:              "02194999400000368626566857200001797430402100",
			expectedBank:         "021",
			expectedCurr:         "9",
			expectedYear:         2025,
			expectedMonth:        time.February,
			expectedDay:          16,
			expectedValue:        36862,
			expectedFree:         "6566857200001797430402100",
			expectedWritableLine: "02196566835720000179074304021004499940000036862",
		},
		{
			name:                 "Valid Boleto 2",
			barcode:              "70792119000000421480001121154691201045051738",
			expectedBank:         "707",
			expectedCurr:         "9",
			expectedYear:         2025,
			expectedMonth:        time.August,
			expectedDay:          31,
			expectedValue:        42148,
			expectedFree:         "0001121154691201045051738",
			expectedWritableLine: "70790001182115469120410450517387211900000042148",
		},
		{
			name:                 "Valid Barcode 2 with 1997 date",
			barcode:              "34196000200000233331098211174108055015849000",
			expectedBank:         "341",
			expectedCurr:         "9",
			expectedYear:         1997,
			expectedMonth:        time.August,
			expectedDay:          9,
			expectedValue:        23333,
			expectedFree:         "1098211174108055015849000",
			expectedWritableLine: "34191098261117410805750158490008600020000023333",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Boleto{validBarcode: tt.barcode}

			if got := b.BankCode(); got != tt.expectedBank {
				t.Errorf("BankCode() = %v, want %v", got, tt.expectedBank)
			}

			if got := b.CurrencyCode(); got != tt.expectedCurr {
				t.Errorf("CurrencyCode() = %v, want %v", got, tt.expectedCurr)
			}

			year, month, day := b.calcExpirationDateAt(time.Date(2025, time.January, 10, 23, 59, 59, 0, brTz))
			if year != tt.expectedYear || month != tt.expectedMonth || day != tt.expectedDay {
				t.Errorf("ExpirationDate() = %v-%v-%v, want %v-%v-%v", year, month, day, tt.expectedYear, tt.expectedMonth, tt.expectedDay)
			}

			year, month, day = b.calcExpirationDateAt(time.Date(2025, time.February, 21, 0, 0, 0, 0, brTz))
			if year != tt.expectedYear || month != tt.expectedMonth || day != tt.expectedDay {
				t.Errorf("ExpirationDate() = %v-%v-%v, want %v-%v-%v", year, month, day, tt.expectedYear, tt.expectedMonth, tt.expectedDay)
			}

			// At 2025-02-22 the expiration date time counter resets
			year, month, day = b.calcExpirationDateAt(time.Date(2025, time.February, 22, 0, 0, 0, 0, brTz))
			if year != tt.expectedYear || month != tt.expectedMonth || day != tt.expectedDay {
				t.Errorf("ExpirationDate() = %v-%v-%v, want %v-%v-%v", year, month, day, tt.expectedYear, tt.expectedMonth, tt.expectedDay)
			}

			year, month, day = b.calcExpirationDateAt(time.Date(2025, time.February, 23, 0, 0, 0, 0, brTz))
			if year != tt.expectedYear || month != tt.expectedMonth || day != tt.expectedDay {
				t.Errorf("ExpirationDate() = %v-%v-%v, want %v-%v-%v", year, month, day, tt.expectedYear, tt.expectedMonth, tt.expectedDay)
			}

			if got := b.Value(); got != tt.expectedValue {
				t.Errorf("Value() = %v, want %v", got, tt.expectedValue)
			}

			if got := b.FreeField(); got != tt.expectedFree {
				t.Errorf("FreeField() = %v, want %v", got, tt.expectedFree)
			}

			if got := b.Barcode(); got != tt.barcode {
				t.Errorf("Barcode() = %v, want %v", got, tt.barcode)
			}

			if got := b.WritableLine(); got != tt.expectedWritableLine {
				t.Errorf("WritableLine() = %v, want %v", got, tt.expectedWritableLine)
			}
		})
	}
}
