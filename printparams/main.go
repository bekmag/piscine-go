package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := range a {
		if i >= 1 {
			b := []rune(a[i])
			for j := range a[i] {
				z01.PrintRune(b[j])
			}
			z01.PrintRune(10)
		}
	}
}
