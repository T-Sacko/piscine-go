package piscine

func IsPrintable(s string) bool {
	str := []byte(s)
	for _, bit := range str {
		if bit < 32 || bit > 126 {
			return false
		}
	}
	return true
}
