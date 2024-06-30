package sequence

func (seq Sequence[E]) Filter(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range seq.Fn {
				if p(e) && !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			i := 0
			for e := range seq.Fn {
				if p(e) {
					if !yield(i, e) {
						return
					}
					i++
				}
			}
		},
	}
}
