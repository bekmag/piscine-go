package main

import "github.com/01-edu/z01"

func main() {
	var aRune rune = '0'
	for aRune <= '9' {

		z01.PrintRune(aRune)
		aRune++
	}
	z01.PrintRune(10)
}
