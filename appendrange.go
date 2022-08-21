package piscine

func AppendRange(min, max int) []int {

	var ranger []int
	for i := min; i < max; i++ {

		ranger = append(ranger, i)
	}
	return ranger
}
