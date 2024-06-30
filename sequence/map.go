package sequence

type MapSeq[E any] struct {
	Sequence[E]
}

func Map[S, T any](seq Iterable[S], f func(S) T) MapSeq[T] {
	return MapSeq[T]{
		Sequence[T]{func(yield func(T) bool) {
			for e := range seq.All().Fn {
				if !yield(f(e)) {
					return
				}
			}
		}},
	}
}
