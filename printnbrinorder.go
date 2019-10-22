package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []int
	if n == 0 {
		arr = append(arr, 0)
	}
	for n != 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	SortIntegerTable(arr)
	for z := range arr {
		z01.PrintRune(rune('0' + arr[z]))
	}
}
