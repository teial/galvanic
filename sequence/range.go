package sequence

import "golang.org/x/exp/constraints"

func RangeTo[E constraints.Integer](start, end E) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for i := start; i < end; i++ {
				if !yield(i) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i := start; i < end; i++ {
				if !yield(int(i)-int(start), i) {
					return
				}
			}
		},
	}
}

func Range[E constraints.Integer](start E) Sequence[E] {
	return Sequence[E]{
		func(yield func(E) bool) {
			for i := start; ; i++ {
				if !yield(i) {
					return
				}
			}
		},
		func(yield func(int, E) bool) {
			for i := start; ; i++ {
				if !yield(int(i)-int(start), i) {
					return
				}
			}
		},
	}
}
