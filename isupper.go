package piscine

func IsUpper(s string) bool {
	str := []byte(s)

	for _, bit := range str {
		if bit <= 65 || bit >= 90 {
			return false
		}
	}
	return true
}
