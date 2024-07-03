package sequence

func (seq Sequence[E]) Chain(seqs ...Iterable[E]) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for e := range seq.Fn {
				if !yield(e) {
					return
				}
			}
			for _, seq := range seqs {
				for e := range seq.All().Fn {
					if !yield(e) {
						return
					}
				}
			}
		},
		func(yield func(int, E) bool) {
			i := 0
			for _, e := range seq.Fn2 {
				if !yield(i, e) {
					return
				}
				i++
			}
			for _, seq := range seqs {
				for _, e := range seq.All().Fn2 {
					if !yield(i, e) {
						return
					}
					i++
				}
			}
		},
	}
}

func (seq ReversibleSequence[E]) Chain(seqs ...ReverseIterable[E]) ReverseIterable[E] {
	end := len(seqs) - 1
	return ReversibleSequence[E]{
		Sequence[E]{
			func(yield func(E) bool) {
				for e := range seq.Fn {
					if !yield(e) {
						return
					}
				}
				for _, seq := range seqs {
					for e := range seq.All().Fn {
						if !yield(e) {
							return
						}
					}
				}
			},
			func(yield func(int, E) bool) {
				i := 0
				for _, e := range seq.Fn2 {
					if !yield(i, e) {
						return
					}
					i++
				}
				for _, seq := range seqs {
					for _, e := range seq.All().Fn2 {
						if !yield(i, e) {
							return
						}
						i++
					}
				}
			},
		},
		func(yield func(E) bool) {
			if end >= 0 {
				for e := range seqs[end].Reverse().Fn {
					if !yield(e) {
						return
					}
				}
				for i := end - 1; i >= 0; i-- {
					for e := range seqs[i].Reverse().Fn {
						if !yield(e) {
							return
						}
					}
				}
			}
			for e := range seq.RevFn {
				if !yield(e) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			i := 0
			if end >= 0 {
				for e := range seqs[end].Reverse().Fn {
					if !yield(i, e) {
						return
					}
					i++
				}
				for i := end - 1; i >= 0; i-- {
					for e := range seqs[i].Reverse().Fn {
						if !yield(i, e) {
							return
						}
					}
					i++
				}
			}
			for e := range seq.RevFn {
				if !yield(i, e) {
					return
				}
				i++
			}
		},
	}
}
