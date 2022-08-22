package piscine

func MakeRange(min, max int) []int {
	var str []int
	if min > max {
		return str
	}
	amount := max - min
	str = make([]int, amount)

	index := 0
	for i := min; i < max; i++ {
		str[index] = i
		index++
	}

	return str
}
