package goleto

import "testing"

func TestIsValidGdaWritableLine(t *testing.T) {
	tests := []struct {
		name          string
		writableLine  string
		expectedValid bool
	}{
		{
			name:          "valid writable line 1",
			writableLine:  "838700000010788500384078011454458202019323676937",
			expectedValid: true,
		},
		{
			name:          "valid writable line 2",
			writableLine:  "817700000000010936599702411310797039001433708318",
			expectedValid: true,
		},
		{
			name:          "invalid writable line with wrong checksum",
			writableLine:  "127456789012345678901234567890123456789012345679",
			expectedValid: false,
		},
		{
			name:          "invalid writable line with incorrect length",
			writableLine:  "12645678901234567890123456789012345678901234567",
			expectedValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g Gda
			if got := g.isValidWritableLine(tt.writableLine); got != tt.expectedValid {
				t.Errorf("isValidGdaWritableLine() = %v, want %v", got, tt.expectedValid)
			}
		})
	}
}
