package piscine

func UltimateDivMod(a *int, b *int) {
	var x *int
	var y *int
	*x = *a
	*y = *b
	*a = *x / *y
	*b = *x % *y
}
