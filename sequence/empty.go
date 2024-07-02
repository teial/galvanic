package sequence

func makeEmpty[E any]() Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			return
		},
		func(yield func(int, E) bool) {
			return
		},
	}
}

func Empty[E any]() ReversibleSequence[E] {
	return ReversibleSequence[E]{
		makeEmpty[E](),
		func(yield func(E) bool) {
			return
		},
		func(yield func(int, E) bool) {
			return
		},
	}
}
