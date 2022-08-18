package piscine

func IsNumeric(s string) bool {
	str := []byte(s)

	for _, bit := range str {
		if bit < 48 || bit > 57 {
			return false
		}
	}
	return true
}
