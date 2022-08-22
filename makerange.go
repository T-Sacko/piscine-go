package piscine

func MakeRange(min, max int) []int {
	var yhh []int
	if min > max {
		return yhh
	}

	str := make([]int, max-min)

	{
		for i := 0; i < max-2; i++ {
			str[i] = i + min
		}
	}
	return str
}
