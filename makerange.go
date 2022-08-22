package piscine

func MakeRange(min, max int) []int {
	var numlist []int
	if min > max {
		return numlist
	}
	amount := max - min
	numlist = make([]int, amount)

	index := 0
	for i := min; i < max; i++ {
		numlist[index] = i
		index++
	}

	return numlist
}
