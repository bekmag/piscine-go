package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	slice := []rune(str)
	for index, word := range slice {
		z01.PrintRune(runes(index))
	}
}
