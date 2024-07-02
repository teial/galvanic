package sequence

func (seq Sequence[E]) Chain(seqs ...Iterable[E]) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range seq.Fn {
				if !yield(e) {
					return
				}
			}
			for _, seq := range seqs {
				for e := range seq.All().Fn {
					if !yield(e) {
						return
					}
				}
			}
		},
		func(yield func(int, E) bool) {
			i := 0
			for _, e := range seq.Fn2 {
				if !yield(i, e) {
					return
				}
				i++
			}
			for _, seq := range seqs {
				for _, e := range seq.All().Fn2 {
					if !yield(i, e) {
						return
					}
					i++
				}
			}
		},
	}
}
