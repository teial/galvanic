package sequence

import "iter"

type Sequence[E any] struct {
	Fn  iter.Seq[E]
	Fn2 iter.Seq2[int, E]
}

type ReversibleSequence[E any] struct {
	Sequence[E]
	RevFn  iter.Seq[E]
	RevFn2 iter.Seq2[int, E]
}

type Iterable[E any] interface {
	All() Sequence[E]
}

func (seq Sequence[E]) All() Sequence[E] {
	return seq
}
