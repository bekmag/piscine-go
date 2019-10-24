package piscine

func MakeRange(min, max int) []int {
	var n []int
	if max-min > 0 {
		a := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			a[i] = i + min
		}
		return a
	}
	return n
}
