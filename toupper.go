package piscine

func ToUpper(s string) string {
	STB := []byte(s)
	diff := 32
	for Index, bit := range STB {
		if bit >= 97 && bit <= 122 {
			Upperbit := int(bit) - diff
			STB[Index] = byte(Upperbit)
		}
	}
	return string(STB)
}
