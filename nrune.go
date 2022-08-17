package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)

	if n > len(runes) {
		return 0
	}
	if n == 0 || n < 0 {
		return 0
	}

	return runes[n-1]
}
