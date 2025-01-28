package goleto

import "unicode"

func digitsOnly(probe string) bool {
	for _, c := range probe {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
