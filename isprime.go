package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	} else {
		limit := Sqrt((nb))
		for i := 2; i <= int(limit); i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}
