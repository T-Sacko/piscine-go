package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := []rune(os.Args[0])
	for _, char := range string(arguments) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
