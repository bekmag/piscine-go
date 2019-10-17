package piscine

func PrintStr(str string) {
	runes := []rune(str)
	for i := range runes {
		PrintRune(runes(i))
	}
}
