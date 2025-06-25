package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func codes() ([]string, error) {
	c := http.Client{}

	r, err := c.Get("https://www.bcb.gov.br/pom/spb/estatistica/port/ParticipantesSTRport.csv")
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	defer r.Body.Close()

	csv, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	return parseCSV(string(csv))
}

func parseCSV(csv string) ([]string, error) {
	lines := strings.Split(csv, "\n")

	if len(lines) < 1 {
		return nil, errors.New("no data found in CSV")
	}

	header := strings.Split(lines[0], ",")
	printableASCII := &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 0x20, Hi: 0x7E, Stride: 1}, // Basic Latin
		},
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), runes.Remove(runes.NotIn(printableASCII)), norm.NFC)

	codeIndex := -1
	for i, h := range header {
		h = strings.TrimSpace(h)
		h = strings.ToLower(h)
		h, _, err := transform.String(t, h)
		if err != nil {
			continue
		}

		if h == "numero_codigo" {
			codeIndex = i
			break
		}
	}

	if codeIndex == -1 {
		return nil, errors.New("no 'Número_Código' column found in CSV")
	}

	var bankCodes []string
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) >= codeIndex+1 {
			code := strings.TrimSpace(parts[codeIndex])
			if len(code) != 3 {
				continue
			}
			_, err := strconv.Atoi(code) // Validate that it's a number
			if err != nil {
				continue
			}
			bankCodes = append(bankCodes, strings.TrimSpace(parts[codeIndex]))
		}
	}

	if len(bankCodes) == 0 {
		return nil, errors.New("no valid bank codes found in CSV")
	}

	return bankCodes, nil
}
