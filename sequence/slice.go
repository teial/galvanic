package sequence

func Slice[E any](slice []E) Sequence[E] {
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
