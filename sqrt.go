package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 2; i <= nb/2; i++ {
		for nb == i*i {
			return i
		}
	}
	return 0
}
