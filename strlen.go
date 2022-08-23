package piscine

func StrLen(s string) int {
	str := []rune(s)
	count := 0
	for i := 1; i <= len(str); i++ {
		count++
	}
	return count
}
