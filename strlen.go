package piscine

import "github.com/01-edu/z01"

func StrLen(str string) int {
	runes := []rune(str)
	for i := range runes int {
		z01.PrintRune(runes[i])
	}
}
