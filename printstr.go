package piscine

func PrintStr(str string) {
	runes := []rune(str)
	for i := range runes {
		z01.PrintRune(rune(i))
	}
}
