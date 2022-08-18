package piscine

func Index(s string, sb string) int {
	mr := []rune(s)
	sbs := []rune(sb)

	for i := 0; i < len(s); i++ {
		  flag := true
			for j := 0; j < len(sb); j++ {
				index := i + j
				if index<len(s){
					if mr[index] != sbs[j]{
						flag = false
					}
				}
				
				
			}
	if flag {return i}
		}