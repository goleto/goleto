package main

import (
	"context"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/goleto/goleto"
	"github.com/google/subcommands"
)

type generateCmd struct {
	n        uint
	detailed bool
}

func (*generateCmd) Name() string     { return "generate" }
func (*generateCmd) Synopsis() string { return "Generate random boletos." }
func (*generateCmd) Usage() string    { return "Usage: goleto generate [-n <how many>] [-detailed]\n\n" }
func (g *generateCmd) SetFlags(f *flag.FlagSet) {
	f.UintVar(&g.n, "n", 1, "number of boletos to generate")
	f.BoolVar(&g.detailed, "detailed", false, "print detailed barcode information")
}

func (g *generateCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	var (
		i, attempts   uint
		attemptsLimit uint = 1000
	)

	for i < g.n {
		if attempts >= attemptsLimit {
			fmt.Printf("error: generate: %v\n", ErrMaxAttemptsReached)
			return subcommands.ExitFailure
		}
		attempts++
		b, err := randomBoleto()
		if err != nil {
			continue
		}
		if g.detailed {
			parseSingle(b.Barcode(), i == g.n-1)
		} else {
			fmt.Printf("%s\n", b.Barcode())
		}
		i++
	}
	return subcommands.ExitSuccess
}

//go:embed banks.tsv
var banks string
var bankCodes []string

var ErrMaxAttemptsReached = errors.New("max attempts reached")

func randomBoleto() (goleto.Boleto, error) {
	return goleto.NewBoleto(
		goleto.WithBankCode(randomBankCode()),
		goleto.WithExpirationDate(randomDate()),
		goleto.WithValue(randomValue()),
		goleto.WithFreeField(randomNumber(25)),
	)
}

func randomBankCode() string {
	return bankCodes[rand.Int31n(int32(len(bankCodes)))]
}

func randomDate() (int, time.Month, int) {
	var r float64
	for {
		r = rand.NormFloat64()*365 + 5
		if r < 0 {
			r = -rand.ExpFloat64() * 10
		}
		if math.Abs(r) < 4500 {
			break
		}
	}
	days := int(math.Round(r))

	return time.Now().AddDate(0, 0, int(days)).Date()
}

func randomValue() uint64 {
	var r float64
	for {
		r = rand.NormFloat64()*500 + 500
		if r > 0 && r < 10e8 {
			return uint64(math.RoundToEven(r * 100))
		}
	}
}

func randomNumber(n int) string {
	const digits = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Int31n(int32(len(digits)))]
	}
	return string(b)
}

func init() {
	for _, line := range strings.Split(banks, "\n") {
		parts := strings.Split(line, "\t")
		if len(parts) > 0 && len(parts[0]) == 3 {
			bankCodes = append(bankCodes, parts[0])
		}
	}
}
