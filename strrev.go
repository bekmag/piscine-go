package piscine

func StrRev(s string) string {
	runes := []rune(s)
	var a rune
	for j := range runes {
	}
	for i := 0; i < (j+1)/2; i++ {
		a = runes[i]
		runes[i] = runes[j-1]
		runes[j-1] = a
	}
	return string(runes)
}
