package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i <= nb+1; i++ {
		if nb < 20 {
			result = result + i
		}
	}
	return result
}
