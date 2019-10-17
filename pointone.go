package piscine

func main() {

	b := 0

	PointOne(&b)
	println(b)
}

func PointOne(n *int) {

	*n = *n + 1
}
