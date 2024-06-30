package sequence

type Slice[E any] []E

type SliceIter[E any] struct {
	Sequence[E]
}

func (s Slice[E]) All() SliceIter[E] {
	return SliceIter[E]{
		Sequence[E]{func(yield func(E) bool) {
			for _, e := range s {
				if !yield(e) {
					return
				}
			}
		}},
	}
}
