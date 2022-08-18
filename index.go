package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {

		B := checkmatch(runes, []rune(toFind), i)
		if B {
			return i
		}
	}
	return -1
}

func checkmatch(runes []rune, toFind []rune, n int) bool {
	h := len([]rune(toFind))
	for w := 0; w < h; w++ {
		if (n+w < len(runes)) && (runes[n+w] == toFind[w]) {
			continue
		} else {
			return false
		}
	}
	return true
}
