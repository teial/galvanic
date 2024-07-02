package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	tests := []struct {
		name string
		seq  Sequence[int]
		n    int
		want []int
	}{
		{
			name: "empty",
			seq:  Empty[int](),
			n:    0,
			want: []int{},
		},
		{
			name: "single",
			seq:  Values(1),
			n:    1,
			want: []int{1},
		},
		{
			name: "single 0",
			seq:  Values(1),
			n:    0,
			want: []int{},
		},
		{
			name: "single 2",
			seq:  Values(1),
			n:    2,
			want: []int{1},
		},
		{
			name: "multiple",
			seq:  Values(1, 2, 3),
			n:    2,
			want: []int{1, 2},
		},
		{
			name: "multiple 0",
			seq:  Values(1, 2, 3),
			n:    0,
			want: []int{},
		},
		{
			name: "multiple 4",
			seq:  Values(1, 2, 3),
			n:    4,
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.Take(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestTake_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range Values(1, 2, 3, 4, 5).Take(3).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
