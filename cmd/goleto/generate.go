package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

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
		b, err := goleto.NewBoleto(goleto.WithRandomFields())
		if err != nil {
			continue
		}
		if g.detailed {
			_ = parseSingle(b.Barcode(), i == g.n-1)
		} else {
			fmt.Printf("%s\n", b.Barcode())
		}
		i++
	}
	return subcommands.ExitSuccess
}

var ErrMaxAttemptsReached = errors.New("max attempts reached")
