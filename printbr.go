package piscine

import "github.com/01-edu/z01"

//PrintNbr prints an int passed in parameter. All possible values of type int have to go through
func PrintNbr(n int) {
	a := '-'
	b := '1'
	c := '2'
	d := '3'
	e := '0'
	if n < 0 {
		z01.PrintRune(a)
		z01.PrintRune(b)
		z01.PrintRune(c)
		z01.PrintRune(d)
	} else if n == 0 {
		z01.PrintRune(e)
	} else {
		z01.PrintRune(b)
		z01.PrintRune(c)
		z01.PrintRune(d)
	}
}
