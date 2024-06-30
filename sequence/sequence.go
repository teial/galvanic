package sequence

import "iter"

type Sequence[E any] struct {
	Fn iter.Seq[E]
}

type Iterable[E any] interface {
	All() Sequence[E]
}
