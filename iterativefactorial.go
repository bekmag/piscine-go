package piscine

func IterativeFactorial(nb int) int {
	result := 1
	resultb := 0
	if nb > 0 && nb < 2147483647 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	}
	if nb > 0 && nb < 2147483647 {
	}
	return resultb
}
