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
		func(yield func(int, E) bool) {
			for i, e := range s {
				if !yield(i, e) {
					return
				}
			}
		},
	}
}
