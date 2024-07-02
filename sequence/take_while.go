package sequence

import "iter"

func makeTakeWhileFn[E any](p func(E) bool, source iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range source {
			if !p(e) || !yield(e) {
				return
			}
		}
	}
}

func makeTakeWhileFn2[E any](p func(E) bool, source iter.Seq2[int, E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, e := range source {
			if !p(e) || !yield(i, e) {
				return
			}
		}
	}
}

func (seq Sequence[E]) TakeWhile(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		makeTakeWhileFn(p, seq.Fn),
		makeTakeWhileFn2(p, seq.Fn2),
	}
}

func (rseq ReversibleSequence[E]) TakeWhile(p func(E) bool) ReversibleSequence[E] {
	return ReversibleSequence[E]{
		rseq.Sequence.TakeWhile(p),
		makeTakeWhileFn(p, rseq.RevFn),
		makeTakeWhileFn2(p, rseq.RevFn2),
	}
}
