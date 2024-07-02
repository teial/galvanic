package sequence

func (rs ReversibleSequence[E]) Reverse() ReversibleSequence[E] {
	return ReversibleSequence[E]{
		Sequence[E]{
			func(yield func(E) bool) {
				for e := range rs.RevFn {
					if !yield(e) {
						return
					}
				}
			},
			func(yield func(int, E) bool) {
				for i, e := range rs.RevFn2 {
					if !yield(i, e) {
						return
					}
				}
			},
		},
		func(yield func(E) bool) {
			for e := range rs.Fn {
				if !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i, e := range rs.Fn2 {
				if !yield(i, e) {
					return
				}
			}
		},
	}
}
