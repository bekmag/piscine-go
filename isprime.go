package piscine

func IsPrime(nb int) bool {
	if nb <= 0 {
		return false
	}
	if nb == 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		for nb%i == 0 {
			return false
		}
	}
	return true
}
