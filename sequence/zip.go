package sequence

import "iter"

type Pair[S, T any] struct {
	V1  S
	OK1 bool
	V2  T
	OK2 bool
}

func NewPair[S, T any](v1 S, v2 T) Pair[S, T] {
	return Pair[S, T]{V1: v1, OK1: true, V2: v2, OK2: true}
}

func Zip[S, T any](s Iterable[S], t Iterable[T]) Sequence[Pair[S, T]] {
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
				if !pair.OK1 || !pair.OK2 || !yield(pair) {
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
				if !pair.OK1 || !pair.OK2 || !yield(i, pair) {
					return
				}
			}
		},
	}
}
