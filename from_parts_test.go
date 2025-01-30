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
