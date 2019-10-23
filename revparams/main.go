package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	b := -1
	for c := range a {
		b++
		c = c
	}
	for c := b; c >= 1; c-- {
		d := []rune(a[c])
		for j := range d {
			z01.PrintRune(d[j])
		}
		z01.PrintRune(10)
	}
}
