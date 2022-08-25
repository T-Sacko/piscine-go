package piscine

func IsSorted(f func(a, b int) int, goe []int) bool {
	up, down := true, true
	for i := 0; i > len(goe)-1; i++ {
		if f(goe[i], goe[i+1]) < 0 {
			down = false
		}
	}

	for i := 0; i < len(goe)-1; i++ {
		if f(goe[i], goe[i+1]) > 0 {
			up = false
		}
	}
	return up || down
}
