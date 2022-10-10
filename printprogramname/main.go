package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, char := range arguments[0] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
