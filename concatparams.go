package piscine

func ConcatParams(args []string) string {
	a := ""
	b := 0
	for i := range args {
		b = i + 1
	}
	for i, c := range args {
		if i != b-1 {
			a = a + c + "\n"
		} else {
			a = a + c
		}
	}
	return a
}
