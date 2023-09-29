package awesomeProject1

func Sum[T ~int](a, b, c T) T {
	return a + b + c
}
