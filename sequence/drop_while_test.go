package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropWhile(t *testing.T) {
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
			name: "all true",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(int) bool { return true },
			want: []int{},
		},
		{
			name: "all false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(int) bool { return false },
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "some true",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i < 3 },
			want: []int{3, 4, 5},
		},
		{
			name: "some false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i > 3 },
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "mixed",
			seq:  Values(2, 4, 1, 3, 6, 5),
			fn:   func(i int) bool { return i%2 == 0 },
			want: []int{1, 3, 6, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.DropWhile(tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestDropWhile_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Values(2, 4, 1, 3, 5).DropWhile(func(e int) bool { return e%2 == 0 })
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
