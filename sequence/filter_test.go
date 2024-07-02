package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Sequence[int]
		fn   func(int) bool
		want []int
	}{
		{
			name: "empty",
			seq:  Empty[int](),
			fn:   func(int) bool { return true },
			want: []int{},
		},
		{
			name: "single",
			seq:  Values(1),
			fn:   func(int) bool { return true },
			want: []int{1},
		},
		{
			name: "single false",
			seq:  Values(1),
			fn:   func(int) bool { return false },
			want: []int{},
		},
		{
			name: "multiple",
			seq:  Values(1, 2, 3),
			fn:   func(int) bool { return true },
			want: []int{1, 2, 3},
		},
		{
			name: "multiple false",
			seq:  Values(1, 2, 3),
			fn:   func(int) bool { return false },
			want: []int{},
		},
		{
			name: "multiple odd",
			seq:  Values(1, 2, 3),
			fn:   func(e int) bool { return e%2 == 1 },
			want: []int{1, 3},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.Filter(tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestFilter_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Values(1, 2, 3, 4, 5).Filter(func(e int) bool { return e%2 == 1 })
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
