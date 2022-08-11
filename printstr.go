package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		iyes := s[i]
		z01.PrintRune(rune(iyes))
	}
}
