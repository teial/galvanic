package sequence

func Empty[E any]() Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			return
		},
		func(yield func(int, E) bool) {
			return
		},
	}
}
