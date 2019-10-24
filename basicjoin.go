package piscine

func BasicJoin(strs []string) string {
	a := ""
	b := 0
	for i := range strs {
		b = i + 1
	}
	for i, c := range strs {
		if i != b {
			a = a + c
		}
	}
	return a
}
