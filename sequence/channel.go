package sequence

type Channel[E any] <-chan E

func (s Channel[E]) All() Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range s {
				if !yield(e) {
					return
				}
			}
		},
	}
}
