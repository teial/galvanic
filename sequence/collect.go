package sequence

func (s Sequence[E]) Collect() []E {
	slice := make([]E, 0)
	for e := range s.Fn {
		slice = append(slice, e)
	}
	return slice
}
