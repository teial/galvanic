package sequence

func makeSequence[E any](slice []E) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for _, e := range slice {
				if !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i, e := range slice {
				if !yield(i, e) {
					return
				}
			}
		},
	}
}

func Slice[E any](slice []E) ReversibleSequence[E] {
	end := len(slice) - 1
	return ReversibleSequence[E]{
		makeSequence(slice),
		func(yield func(E) bool) {
			for i := 0; i <= end; i++ {
				if !yield(slice[end-i]) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i := 0; i <= end; i++ {
				if !yield(i, slice[end-i]) {
					return
				}
			}
		},
	}
}
