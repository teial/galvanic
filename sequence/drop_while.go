package sequence

func (seq Sequence[E]) DropWhile(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			dropOver := false
			for e := range seq.Fn {
				dropOver = dropOver || !p(e)
				if dropOver && !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			dropOver := false
			i := 0
			for e := range seq.Fn {
				dropOver = dropOver || !p(e)
				if dropOver {
					if !yield(i, e) {
						return
					}
					i++
				}
			}
		},
	}
}
