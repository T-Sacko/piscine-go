package piscine

func IterativeFactorial(index int) int {
	if index < 1 {
		return 0
	}
	if index > 27 {
		return 0
	}
	result := 1

	for i := 1; i < index+1; i++ {
		result = result * i
	}
	return result
}
