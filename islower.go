package piscine

func IsLower(s string) bool {
	str := []byte(s)

	for _, bit := range str {
		if bit < 97 || bit > 122 {
			return false
		}
	}
	return true
}
