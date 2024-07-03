package sequence

import "iter"

func makeMapFn[S, T any](f func(S) T, source iter.Seq[S]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range source {
			if !yield(f(e)) {
				return
			}
		}
	}
}

func makeMapFn2[S, T any](f func(S) T, source iter.Seq2[int, S]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, e := range source {
			if !yield(i, f(e)) {
				return
			}
		}
	}
}

func Map[S, T any](seq Iterable[S], f func(S) T) Sequence[T] {
	return Sequence[T]{
		makeMapFn(f, seq.All().Fn),
		makeMapFn2(f, seq.All().Fn2),
	}
}

func MapR[S, T any](seq ReverseIterable[S], f func(S) T) ReversibleSequence[T] {
	return ReversibleSequence[T]{
		Map[S, T](seq, f),
		makeMapFn(f, seq.Reverse().Fn),
		makeMapFn2(f, seq.Reverse().Fn2),
	}
}
