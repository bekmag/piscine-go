package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	n := 'T'
	b := 'F'
	if nb < 0 {
		z01.PrintRune(n)
	} else {
		z01.PrintRune(b)
	}
	z01.PrintRune(10)
}
