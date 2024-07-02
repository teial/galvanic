package sequence

func Channel[E any](ch <-chan E) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range ch {
				if !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			i := 0
			for e := range ch {
				if !yield(i, e) {
					return
				}
				i++
			}
		},
	}
}
