package piscine

func AppendRange(min, max int) []int {
	var a []int
	for i := 0; i < max; i++ {
		if i >= min {
			a = append(a, i)
		}
	}
	return a
}
