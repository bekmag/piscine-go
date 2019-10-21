package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if nb > 0 && power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return 0
}
