package awesomeProject1

func Sum[T ~int](a, b, c, d T) T {
	return a + b + c + d
}
