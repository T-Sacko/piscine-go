package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setpoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setpoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Runerune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Runerune(points.y)
	z01.PrintRune('\n')
}

func checking(r int) {
	c := '0'
	if r == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= r%10; i++ {
		c++
	}
	for i := 1; i >= r%10; i-- {
		c++
	}
	if r/10 != 0 {
		checking(r / 10)
	}
	z01.PrintRune(c)
	return
}

func Runerune(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	checking(n)
}
