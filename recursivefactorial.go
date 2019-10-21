package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 && nb < 2147483647 {
		return 1
	}
	if nb > 0 && nb < 2147483647 {
		return nb * RecursiveFactorial(nb-1)
	}
	if nb == 0 {
		return 1
	}
	return 0
}
