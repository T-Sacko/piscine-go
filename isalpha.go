package piscine

func IsAlpha(s string) bool {
	str := []byte(s)
	var result bool

	for _, bit := range str {
		if (bit >= 97 && bit <= 122) || (bit >= 65 && bit <= 90) || (bit >= 48 && bit <= 57) {
			result = true
		} else {
			return false
		}
	}
	return result
}
