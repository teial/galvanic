package sequence

type Slice[E any] []E

type SliceSeq[E any] struct {
	Sequence[E]
}

func (s Slice[E]) All() SliceSeq[E] {
	return SliceSeq[E]{
		Sequence[E]{func(yield func(E) bool) {
			for _, e := range s {
				if !yield(e) {
					return
				}
			}
		}},
	}
}
