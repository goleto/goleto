package goleto

import (
	"testing"
)

func TestGda(t *testing.T) {
	tests := []struct {
		name                 string
		barcode              string
		expectedSegmentId    string
		expectedCompanyId    string
		expectedValueType    GdaValueType
		expectedValue        uint64
		expectedFreeField    string
		expectedWritableLine string
	}{
		{
			name:                 "GDA 1",
			barcode:              "83870000001788500384070114544582001932367693",
			expectedSegmentId:    "3",
			expectedCompanyId:    "0038",
			expectedValueType:    GdaEffectiveValue,
			expectedValue:        17885,
			expectedFreeField:    "4070114544582001932367693",
			expectedWritableLine: "838700000010788500384078011454458202019323676937",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gda{validBarcode: tt.barcode}

			if got := g.SegmentId(); got != tt.expectedSegmentId {
				t.Errorf("SegmentId() = %v, want %v", got, tt.expectedSegmentId)
			}

			if got := g.CompanyId(); got != tt.expectedCompanyId {
				t.Errorf("CompanyId() = %v, want %v", got, tt.expectedCompanyId)
			}

			if got := g.ValueType(); got != tt.expectedValueType {
				t.Errorf("ValueType() = %v, want %v", got, tt.expectedValueType)
			}

			if got := g.Value(); got != tt.expectedValue {
				t.Errorf("Value() = %v, want %v", got, tt.expectedValue)
			}

			if got := g.FreeField(); got != tt.expectedFreeField {
				t.Errorf("FreeField() = %v, want %v", got, tt.expectedFreeField)
			}

			if got := g.Barcode(); got != tt.barcode {
				t.Errorf("Barcode() = %v, want %v", got, tt.barcode)
			}

			if got := g.WritableLine(); got != tt.expectedWritableLine {
				t.Errorf("WritableLine() = %v, want %v", got, tt.expectedWritableLine)
			}
		})
	}
}
