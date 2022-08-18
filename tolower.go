package piscine

func ToLower(s string) string {
	STB := []byte(s)
	diff := 32
	for Index, bit := range STB {
		if bit >= 65 && bit <= 90 {
			Upperbit := int(bit) + diff
			STB[Index] = byte(Upperbit)
		}
	}
	return string(STB)
}
