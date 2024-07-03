package sequence

import "iter"

func makeFilterFn[E any](p func(E) bool, source iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range source {
			if p(e) && !yield(e) {
				return
			}
		}
	}
}

func makeFilterFn2[E any](p func(E) bool, source iter.Seq[E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		i := 0
		for e := range source {
			if p(e) {
				if !yield(i, e) {
					return
				}
				i++
			}
		}
	}
}

func (seq Sequence[E]) Filter(p func(E) bool) Sequence[E] {
	return Sequence[E]{
		makeFilterFn(p, seq.Fn),
		makeFilterFn2(p, seq.Fn),
	}
}

func (rseq ReversibleSequence[E]) Filter(p func(E) bool) ReversibleSequence[E] {
	return ReversibleSequence[E]{
		rseq.Sequence.Filter(p),
		makeFilterFn(p, rseq.RevFn),
		makeFilterFn2(p, rseq.RevFn),
	}
}
