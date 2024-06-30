package sequence

type FilterSeq[E any] struct {
	Sequence[E]
}

func (seq Sequence[E]) Filter(p func(E) bool) FilterSeq[E] {
	return FilterSeq[E]{
		Sequence[E]{func(yield func(E) bool) {
			for e := range seq.Fn {
				if p(e) && !yield(e) {
					return
				}
			}
		}},
	}
}
