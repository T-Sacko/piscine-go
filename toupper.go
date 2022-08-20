package piscine

func ToUpper(s string) string {
	STB := []byte(s)

	for Index, bit := range STB {
		if bit >= 97 && bit <= 122 {
			Upperbit := bit - 32
			STB[Index] = Upperbit
		}
	}
	return string(STB)
}
