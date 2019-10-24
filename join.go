package piscine

func Join(strs []string, sep string) string {
	a := ""
	b := 0
	for i := range strs {
		b = i + 1
	}
	for i, c := range strs {
		if i != b-1 {
			a = a + c + sep
		} else {
			a = a + c
		}
	}
	return a
}
