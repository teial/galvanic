package sequence

func FromValues[E any](values ...E) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for _, e := range values {
				if !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i, e := range values {
				if !yield(i, e) {
					return
				}
			}
		},
	}
}
