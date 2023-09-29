package awesomeProject1

func Sum[T ~int](a, b T) T {
	return a + b
}
