package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	for a := range runes {
		if (runes[a] >= 97 && runes[a] <= 122) || (runes[a] >= 65 && runes[a] <= 90) || (runes[a] >= 33 && runes[a] <= 47) || (runes[a] >= 58 && runes[a] <= 64) || (runes[a] >= 91 && runes[a] <= 96) || (runes[a] >= 123 && runes[a] <= 126) || runes[a] == ' ' {
			return false
		}
	}
	return true
}
