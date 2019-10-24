package piscine

func MakeRange(min, max int) []int {
	a := make([]int, max)
	for i := min; i < max; i++ {
		a[i] = i
	}
	return a
}
