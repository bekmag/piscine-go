package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for a := range runes {
		if runes[a] >= 0 && runes[a] <= 31 {
			return false
		}
	}
	return true
}
