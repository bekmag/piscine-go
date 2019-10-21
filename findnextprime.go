package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	}
	if nb == 1 {
		return 2
	}
	for i := 2; i <= nb/2; i++ {
		for nb%i == 0 {
			return nb + 1
		}
	}
	return nb
}
