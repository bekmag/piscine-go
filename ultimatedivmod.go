package piscine

func UltimateDivMod(a *int, b *int) {
	var x, y *int
	*a = *x / *y
	*b = *x % *y
	*x = *a
	*y = *b
}
