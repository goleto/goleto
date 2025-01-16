package goleto

import "strconv"

type GdaValueType string

var (
	GdaReferenceValue GdaValueType = "reference"
	GdaEffectiveValue GdaValueType = "effective"
)

type Gda struct {
	validBarcode string
}

// Barcode returns the Boleto on bar code format.
func (g Gda) Barcode() string {
	return g.validBarcode
}

// WritableLine returns the Boleto on writable line format.
func (g Gda) WritableLine() string {
	return g.barcodeToWritableLine(g.validBarcode)
}

// SegmentId returns segment identifier of this GDA.
// The possible identifiers are:
//
//	"1" - Prefeituras
//	"2" - Saneamento
//	"3" - Energia Elétrica e Gás
//	"4" - Telecomunicações
//	"5" - Órgãos Governamentais
//	"6" - Carnes e Assemelhados ou demais Empresas/Órgãos que serão
//	      identificadas através do CNPJ
//	"7" - Multas de trânsito
//	"9" - Uso exclusivo do banco
func (g Gda) SegmentId() string {
	return g.validBarcode[1:2]
}

// CompanyId extracts and returns the company identifier from the GDA.
// If the segment identifier is "6", it returns the CNPJ.
// Otherwise, it returns a shorter identifier that is defined by Febraban.
func (g Gda) CompanyId() string {
	if g.validBarcode[1] == '6' {
		// using CNPJ to identify the company
		return g.validBarcode[15:23]
	}
	return g.validBarcode[15:19]
}

// ValueType returns whether the value in this GDA is effective or a reference.
// An "effective" value means that the value is in reais (BRL) currency
// A "reference" value means that the value should be multiplied by reference index
func (g Gda) ValueType() GdaValueType {
	switch g.validBarcode[2] {
	case '7', '9':
		return GdaReferenceValue
	default:
		return GdaEffectiveValue
	}
}

// Value returns the value ot this GDA, which should be interpreted according its value type
func (g Gda) Value() uint64 {
	val, _ := strconv.ParseUint(g.validBarcode[4:15], 10, 64)
	return val
}

// FreeField extracts and returns a specific portion of the valid barcode
// from the GDA struct as a string. The portion extracted has 21 characters length
// if the segment identifier is "6", otherwise, it has 25.
func (g Gda) FreeField() string {
	if g.validBarcode[1] == '6' {
		// identificados através do CNPJ,
		return g.validBarcode[23:44]
	}
	return g.validBarcode[19:44]
}
