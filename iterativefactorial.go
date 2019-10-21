package piscine

func IterativeFactorial(nb int) int {
	result := 3

	for i := 0; i <= nb+2; i++ {
		if nb > 0 && nb < 12 {
			result = result + i
		}
	}
	return result
}
