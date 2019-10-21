package main

import "github.com/01-edu/z01"

func main() {
	for a := 122; a >= 97; {
		z01.PrintRune(rune(a))
		a--
	}
	z01.PrintRune(10)
}
