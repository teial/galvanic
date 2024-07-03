package sequence

import "iter"

// Sequence is an iterator that can only be traversed forward.
type Sequence[E any] struct {
	Fn  iter.Seq[E]
	Fn2 iter.Seq2[int, E]
}

type Iterable[E any] interface {
	All() Sequence[E]
}

func (seq Sequence[E]) All() Sequence[E] {
	return seq
}

// ReversibleSequence is an iterator that can be traversed both forward and backward.
type ReversibleSequence[E any] struct {
	Sequence[E]
	RevFn  iter.Seq[E]
	RevFn2 iter.Seq2[int, E]
}

type ReverseIterable[E any] interface {
	Iterable[E]
	Reverse() ReversibleSequence[E]
}

func (rseq ReversibleSequence[E]) Reverse() ReversibleSequence[E] {
	return ReversibleSequence[E]{
		Sequence[E]{rseq.RevFn, rseq.RevFn2},
		rseq.Fn,
		rseq.Fn2,
	}
}
