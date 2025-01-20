package goleto

import (
	"testing"
)

func TestParseBoleto(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "Valid barcode",
			input:   "02194999400000368626566857200001797430402100",
			wantErr: nil,
		},
		{
			name:    "Valid barcode 2",
			input:   "23791115200000593594150093655089062707123000",
			wantErr: nil,
		},
		{
			name:    "Valid barcode with date factor less then 1000",
			input:   "34196000200000233331098211174108055015849000",
			wantErr: nil,
		},
		{
			name:    "Valid writable line",
			input:   "02196566835720000179074304021004499940000036862",
			wantErr: nil,
		},
		{
			name:    "Valid writable line 2",
			input:   "23794150099365508906327071230000111520000059359",
			wantErr: nil,
		},
		{
			name:    "Invalid characters",
			input:   "12345abc901234567890123456789012345678901234",
			wantErr: ErrInvalidCode,
		},
		{
			name:    "Invalid length",
			input:   "1234567890123456789012345678901234567890123",
			wantErr: ErrInvalidCode,
		},
		{
			name:    "Invalid writable line",
			input:   "12345678901234567890123456789012345678901234568",
			wantErr: ErrInvalidCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseBoleto(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParseBoleto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestParseGda(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "Valid GDA barcode",
			input:   "83870000001788500384070114544582001932367693",
			wantErr: nil,
		},
		{
			name:    "Valid GDA barcode 2",
			input:   "81770000000010936599704113107970300143370831",
			wantErr: nil,
		},
		{
			name:    "Valid GDA writable line",
			input:   "838700000010788500384078011454458202019323676937",
			wantErr: nil,
		},
		{
			name:    "Valid GDA writable line 2",
			input:   "817700000000010936599702411310797039001433708318",
			wantErr: nil,
		},
		{
			name:    "Invalid characters",
			input:   "83660abc0001503080048100914041160600000000000",
			wantErr: ErrInvalidCode,
		},
		{
			name:    "Invalid length",
			input:   "8366000000150308004810091404116060000000000",
			wantErr: ErrInvalidCode,
		},
		{
			name:    "Invalid GDA writable line",
			input:   "836600000015030800481009140411606000000000000000",
			wantErr: ErrInvalidCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseGda(tt.input)
			if err != tt.wantErr {
				t.Errorf("ParseGda() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
