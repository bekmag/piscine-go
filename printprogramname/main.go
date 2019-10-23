package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	runes := []rune(a[0])
	for i := range runes {
		z01.PrintRune(runes[i])
	}
}
