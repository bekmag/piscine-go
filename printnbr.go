package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	a := []rune(n)
	b := range(a)
	if n < 0 {
		z01.PrintRune(b)
	} else if n == 0 {
		z01.PrintRune(b)
	} else {
		z01.PrintRune(b)
	}
}
