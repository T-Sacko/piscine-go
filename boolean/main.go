package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	arguments := os.Args
	lengthOfArg := len(arguments) + 1

	if isEven(lengthOfArg) {
		printStr("even thing")
	} else {
		printStr("odd thing")
	}
}
