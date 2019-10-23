package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] > 96 && runes[i] < 122 {
			runes[i] = runes[i] - 32
		}
	}
	str := string(runes)
	return str
}
