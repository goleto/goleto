# goleto

Goleto is a library to parse strings into valid boletos or GDAs (Guia de Arrecadação).

A valid boleto string is a sequency of 44 or 47 characters that obey the standard defined on [section 2.3 of Banco do Brasil's boleto de pagamento spec](https://www.bb.com.br/docs/pub/emp/empl/dwn/Doc5175Bloqueto.pdf).

A valid GDA string is a sequency of 44 or 48 characters that obey the standard defined on [“Layout” Padrão de Arrecadação/Recebimento com Utilização de Código de Barras](https://cmsarquivos.febraban.org.br/Arquivos/documentos/PDF/Layout%20-%20C%C3%B3digo%20de%20Barras%20-%20Vers%C3%A3o%207%20-%2001_03_2023_mn.pdf)

# Examples

**Using `ParseBoleto()` function**

```go
package main

import (
    "fmt"

    "github.com/goleto/goleto"
)

func main() {
    barcode := "02194999400000368626566857200001797430402100"
    boleto, err := goleto.ParseBoleto(barcode)
    if err != nil {
        fmt.Println("Error parsing boleto:", err)
        return
    }

    fmt.Println("Barcode:", boleto.Barcode())
    fmt.Println("Writable Line:", boleto.WritableLine())
    fmt.Println("Bank Code:", boleto.BankCode())
    fmt.Println("Currency Code:", boleto.CurrencyCode())
    fmt.Println("Expiration Date:", boleto.ExpirationDate())
    fmt.Println("Value:", boleto.Value())
    fmt.Println("Free Field:", boleto.FreeField())
}
```

**Using `ParseGda()` function**

```go
package main

import (
    "fmt"

    "github.com/goleto/goleto"
)

func main() {
    barcode := "817700000000010936599702411310797039001433708318"
    gda, err := goleto.ParseGda(barcode)
    if err != nil {
        fmt.Println("Error parsing GDA:", err)
        return
    }

    fmt.Println("Barcode:", gda.Barcode())
    fmt.Println("Writable Line:", gda.WritableLine())
    fmt.Println("Segment ID:", gda.SegmentId())
    fmt.Println("Company ID:", gda.CompanyId())
    fmt.Println("Value Type:", gda.ValueType())
    fmt.Println("Value:", gda.Value())
    fmt.Println("Free Field:", gda.FreeField())
}
```
