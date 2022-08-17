package piscine

func AlphaCount(s string) int {
	count := 0
	str := []byte(s)
	for _, bit := range str {
		if (bit >= 65 && bit <= 90) || (bit >= 97 && bit <= 122) {
			count++
		}
	}

	return count
}
