package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 && nb < 12 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}
