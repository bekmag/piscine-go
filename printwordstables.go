package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	count := 0
	for range table {
		count++
	}
	for _, word := range table {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
