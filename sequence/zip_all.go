package sequence

import "iter"

func makeZipAllFn[S, T any](s iter.Seq[S], t iter.Seq[T]) iter.Seq[Pair[S, T]] {
	return func(yield func(Pair[S, T]) bool) {
		pull1, stop := iter.Pull(s)
		defer stop()
		pull2, stop := iter.Pull(t)
		defer stop()
		for {
			var pair Pair[S, T]
			pair.V1, pair.OK1 = pull1()
			pair.V2, pair.OK2 = pull2()
			if (!pair.OK1 && !pair.OK2) || !yield(pair) {
				return
			}
		}
	}
}

func makeZipAllFn2[S, T any](s iter.Seq[S], t iter.Seq[T]) iter.Seq2[int, Pair[S, T]] {
	return func(yield func(int, Pair[S, T]) bool) {
		pull1, stop := iter.Pull(s)
		defer stop()
		pull2, stop := iter.Pull(t)
		defer stop()
		for i := 0; ; i++ {
			var pair Pair[S, T]
			pair.V1, pair.OK1 = pull1()
			pair.V2, pair.OK2 = pull2()
			if (!pair.OK1 && !pair.OK2) || !yield(i, pair) {
				return
			}
		}
	}
}

func ZipAll[S, T any](s Iterable[S], t Iterable[T]) Sequence[Pair[S, T]] {
	return Sequence[Pair[S, T]]{
		makeZipAllFn(s.All().Fn, t.All().Fn),
		makeZipAllFn2(s.All().Fn, t.All().Fn),
	}
}

func ZipAllR[S, T any](s ReverseIterable[S], t ReverseIterable[T]) ReversibleSequence[Pair[S, T]] {
	return ReversibleSequence[Pair[S, T]]{
		ZipAll[S, T](s, t),
		makeZipAllFn(s.Reverse().Fn, t.Reverse().Fn),
		makeZipAllFn2(s.Reverse().Fn, t.Reverse().Fn),
	}
}
