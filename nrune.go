package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	}
	runes := []rune(s)
	a := 0
	for index, i := range runes {
		a++
		i = i
		index = index
	}
	return runes[n-1]
}
