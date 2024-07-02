package sequence

func Values[E any](values ...E) ReversibleSequence[E] {
	end := len(values) - 1
	return ReversibleSequence[E]{
		makeSequence(values),
		func(yield func(E) bool) {
			for i := 0; i <= end; i++ {
				if !yield(values[end-i]) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i := 0; i <= end; i++ {
				if !yield(i, values[end-i]) {
					return
				}
			}
		},
	}
}
