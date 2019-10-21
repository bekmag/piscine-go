package piscine

func RecursivePower(nb int, power int) int {
	if nb < 0 && power < 0 {
		return 0
	}
	if nb == 1 && power == 1 {
		return 1
	}
	if nb == 0 && power == 0 {
		return 1
	}
	if nb > 0 && power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return power + 1
}
