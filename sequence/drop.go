package sequence

import "iter"

func makeDropFn[E any](n int, source iter.Seq2[int, E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i, e := range source {
			if i >= n && !yield(e) {
				return
			}
		}
	}
}

func makeDropFn2[E any](n int, source iter.Seq2[int, E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, e := range source {
			if i >= n && !yield(i-n, e) {
				return
			}
		}
	}
}

func (seq Sequence[E]) Drop(n int) Sequence[E] {
	if n < 0 {
		panic("n must be non-negative")
	}
	return Sequence[E]{
		makeDropFn(n, seq.Fn2),
		makeDropFn2(n, seq.Fn2),
	}
}

func (rseq ReversibleSequence[E]) Drop(n int) ReversibleSequence[E] {
	return ReversibleSequence[E]{
		rseq.Sequence.Drop(n),
		makeDropFn(n, rseq.RevFn2),
		makeDropFn2(n, rseq.RevFn2),
	}
}
