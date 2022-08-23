package piscine

func MakeRange(min, max int) []int {
	var yhh []int
	if min > max || min == max {
		return yhh
	}
	str := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		str[i] = i + min
	}

	return str
}
