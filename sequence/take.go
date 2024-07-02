package sequence

import "iter"

func makeTakeFn[E any](n int, source iter.Seq2[int, E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i, e := range source {
			if i >= n || !yield(e) {
				return
			}
		}
	}
}

func makeTakeFn2[E any](n int, source iter.Seq2[int, E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, e := range source {
			if i >= n || !yield(i, e) {
				return
			}
		}
	}
}

func (seq Sequence[E]) Take(n int) Sequence[E] {
	if n < 0 {
		panic("n must be non-negative")
	}
	return Sequence[E]{
		makeTakeFn(n, seq.Fn2),
		makeTakeFn2(n, seq.Fn2),
	}
}

func (rseq ReversibleSequence[E]) Take(n int) ReversibleSequence[E] {
	return ReversibleSequence[E]{
		rseq.Sequence.Take(n),
		makeTakeFn(n, rseq.RevFn2),
		makeTakeFn2(n, rseq.RevFn2),
	}
}
