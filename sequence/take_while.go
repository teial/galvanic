package sequence

func (seq Sequence[E]) TakeWhile(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range seq.Fn {
				if !p(e) || !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i, e := range seq.Fn2 {
				if !p(e) || !yield(i, e) {
					return
				}
			}
		},
	}
}
