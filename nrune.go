package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return '\x00'
	}
	runes := []rune(s)
	a := 0
	for index, i := range runes {
		a++
		i = i
		index = index
	}
	if n > a {
		return '\x00'
	}
	return runes[n-1]
}
