package sequence

func (seq Sequence[E]) Drop(n int) Sequence[E] {
	if n < 0 {
		panic("n must be non-negative")
	}
	return Sequence[E]{
		func(yield func(E) bool) {
			for i, e := range seq.Fn2 {
				if i >= n && !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i, e := range seq.Fn2 {
				if i >= n && !yield(i-n, e) {
					return
				}
			}
		},
	}
}
