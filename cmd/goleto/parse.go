package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/goleto/goleto"
	"github.com/google/subcommands"
)

type parseCmd struct{}

func (*parseCmd) Name() string             { return "parse" }
func (*parseCmd) Synopsis() string         { return "Parse barcode displaying their details." }
func (*parseCmd) Usage() string            { return "Usage: goleto parse [BARCODE]...\n\n" }
func (*parseCmd) SetFlags(_ *flag.FlagSet) {}

func (*parseCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	notDigits := regexp.MustCompile("[^0-9]+")
	var errs []error

	barcodes, ioErr := readStdin()
	barcodes = append(barcodes, f.Args()...)

	l := len(barcodes)
	for i, b := range barcodes {
		barcode := notDigits.ReplaceAll([]byte(b), []byte{})

		err := parseSingle(string(barcode), i == l-1)
		if err != nil {
			errs = append(errs, fmt.Errorf("fail parsing %s", b))
			continue
		}
	}

	if err := errors.Join(append(errs, ioErr)...); err != nil {
		fmt.Printf("error: parse: %v\n", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func parseSingle(barcode string, last bool) error {
	if err := printDetails(barcode); err != nil {
		return err
	}

	if !last {
		fmt.Println()
	}
	return nil
}

func readStdin() (lines []string, err error) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if e := s.Err(); e != nil {
		err = fmt.Errorf("reading barcodes: %w", e)
	}
	return
}

func printDetails(barcode string) error {
	if barcode[0] == '8' {
		gda, err := goleto.ParseGda(barcode)
		if err != nil {
			return err
		}
		printGda(gda)
		return nil
	}

	boleto, err := goleto.ParseBoleto(barcode)
	if err != nil {
		return err
	}
	printBoleto(boleto)
	return nil
}

func printGda(gda goleto.Gda) {
	fmt.Println("Barcode:\t\t", gda.Barcode())
	fmt.Println("Writable Line:\t\t", gda.WritableLine())
	fmt.Println("Segment ID:\t\t", gda.SegmentId())
	fmt.Println("Company ID:\t\t", gda.CompanyId())
	fmt.Println("Value Type:\t\t", gda.ValueType())
	fmt.Println("Value:\t\t\t", gda.Value())
	fmt.Println("Free Field:\t\t", gda.FreeField())
}

func printBoleto(boleto goleto.Boleto) {
	fmt.Println("Barcode:\t\t", boleto.Barcode())
	fmt.Println("Writable Line:\t\t", boleto.WritableLine())
	fmt.Println("Bank Code:\t\t", boleto.BankCode())
	fmt.Println("Currency Code:\t\t", boleto.CurrencyCode())
	fmt.Println("Date factor:\t\t", boleto.DateFactor())
	y, m, d := boleto.ExpirationDate()
	fmt.Printf("Expiration Date:\t %d-%02d-%02d\n", y, m, d)
	fmt.Println("Value:\t\t\t", boleto.Value())
	fmt.Println("Free Field:\t\t", boleto.FreeField())
}
