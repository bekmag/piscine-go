package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = b / a
	*mod = a % b
}