package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	a := 0
	for i := range runes {
		a++
		i = i
	}
	return runes[a-1]
}
