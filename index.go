package piscine

func Index(s string, toFind string) int {
	mr := []rune(s)
	sbs := []rune(toFind)

	for i := 0; i < len(mr); i++ {
		if mr[i] == sbs[0] {
			return i
		}
	}
	return -1
}