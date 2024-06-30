package sequence

type Slice[E any] []E

func (s Slice[E]) All() Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for _, e := range s {
				if !yield(e) {
					return
				}
			}
		},
	}
}
