package sequence

import "iter"

func makeDropWhileFn[E any](p func(E) bool, source iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		dropOver := false
		for e := range source {
			dropOver = dropOver || !p(e)
			if dropOver && !yield(e) {
				return
			}
		}
	}
}

func makeDropWhileFn2[E any](p func(E) bool, source iter.Seq[E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		dropOver := false
		i := 0
		for e := range source {
			dropOver = dropOver || !p(e)
			if dropOver {
				if !yield(i, e) {
					return
				}
				i++
			}
		}
	}
}

func (seq Sequence[E]) DropWhile(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		makeDropWhileFn(p, seq.Fn),
		makeDropWhileFn2(p, seq.Fn),
	}
}

func (rseq ReversibleSequence[E]) DropWhile(p func(E) bool) ReversibleSequence[E] {
	return ReversibleSequence[E]{
		rseq.Sequence.DropWhile(p),
		makeDropWhileFn(p, rseq.RevFn),
		makeDropWhileFn2(p, rseq.RevFn),
	}
}
