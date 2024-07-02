package sequence

import "iter"

func ZipAll[S, T any](s Iterable[S], t Iterable[T]) Sequence[Pair[S, T]] {
	return Sequence[Pair[S, T]]{
		func(yield func(Pair[S, T]) bool) {
			pull1, stop := iter.Pull(s.All().Fn)
			defer stop()
			pull2, stop := iter.Pull(t.All().Fn)
			defer stop()
			for {
				var pair Pair[S, T]
				pair.V1, pair.OK1 = pull1()
				pair.V2, pair.OK2 = pull2()
				if (!pair.OK1 && !pair.OK2) || !yield(pair) {
					return
				}
			}
		},
		func(yield func(int, Pair[S, T]) bool) {
			pull1, stop := iter.Pull(s.All().Fn)
			defer stop()
			pull2, stop := iter.Pull(t.All().Fn)
			defer stop()
			for i := 0; ; i++ {
				var pair Pair[S, T]
				pair.V1, pair.OK1 = pull1()
				pair.V2, pair.OK2 = pull2()
				if (!pair.OK1 && !pair.OK2) || !yield(i, pair) {
					return
				}
			}
		},
	}
}
