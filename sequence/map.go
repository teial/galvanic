package sequence

func Map[S, T any](seq Iterable[S], f func(S) T) Sequence[T] {
	return Sequence[T]{
		func(yield func(T) bool) {
			for e := range seq.All().Fn {
				if !yield(f(e)) {
					return
				}
			}
		},
		func(yield func(int, T) bool) {
			for i, e := range seq.All().Fn2 {
				if !yield(i, f(e)) {
					return
				}
			}
		},
	}
}
