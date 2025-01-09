package goleto

import (
	"testing"
)

func TestParse(t *testing.T) {
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
			wantErr: ErrInvalidBarcode,
		},
		{
			name:    "Invalid length",
			input:   "1234567890123456789012345678901234567890123",
			wantErr: ErrInvalidBarcode,
		},
		{
			name:    "Invalid writable line",
			input:   "12345678901234567890123456789012345678901234568",
			wantErr: ErrInvalidBarcode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if err != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
