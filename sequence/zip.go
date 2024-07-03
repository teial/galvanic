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

func makeZipFn[S, T any](s iter.Seq[S], t iter.Seq[T]) iter.Seq[Pair[S, T]] {
	return func(yield func(Pair[S, T]) bool) {
		pull1, stop := iter.Pull(s)
		defer stop()
		pull2, stop := iter.Pull(t)
		defer stop()
		for {
			var pair Pair[S, T]
			pair.V1, pair.OK1 = pull1()
			pair.V2, pair.OK2 = pull2()
			if !pair.OK1 || !pair.OK2 || !yield(pair) {
				return
			}
		}
	}
}

func makeZipFn2[S, T any](s iter.Seq[S], t iter.Seq[T]) iter.Seq2[int, Pair[S, T]] {
	return func(yield func(int, Pair[S, T]) bool) {
		pull1, stop := iter.Pull(s)
		defer stop()
		pull2, stop := iter.Pull(t)
		defer stop()
		for i := 0; ; i++ {
			var pair Pair[S, T]
			pair.V1, pair.OK1 = pull1()
			pair.V2, pair.OK2 = pull2()
			if !pair.OK1 || !pair.OK2 || !yield(i, pair) {
				return
			}
		}
	}
}

func Zip[S, T any](s Iterable[S], t Iterable[T]) Sequence[Pair[S, T]] {
	return Sequence[Pair[S, T]]{
		makeZipFn(s.All().Fn, t.All().Fn),
		makeZipFn2(s.All().Fn, t.All().Fn),
	}
}

func ZipR[S, T any](s ReverseIterable[S], t ReverseIterable[T]) ReversibleSequence[Pair[S, T]] {
	return ReversibleSequence[Pair[S, T]]{
		Zip[S, T](s, t),
		makeZipFn(s.Reverse().Fn, t.Reverse().Fn),
		makeZipFn2(s.Reverse().Fn, t.Reverse().Fn),
	}
}
