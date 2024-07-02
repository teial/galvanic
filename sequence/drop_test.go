package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "empty",
			slice: []int{},
			n:     0,
			want:  []int{},
		},
		{
			name:  "single",
			slice: []int{1},
			n:     1,
			want:  []int{},
		},
		{
			name:  "single 0",
			slice: []int{1},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "single 2",
			slice: []int{1},
			n:     2,
			want:  []int{},
		},
		{
			name:  "multiple",
			slice: []int{1, 2, 3},
			n:     2,
			want:  []int{3},
		},
		{
			name:  "multiple 0",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "multiple 4",
			slice: []int{1, 2, 3},
			n:     4,
			want:  []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := FromSlice(tc.slice).Drop(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestDrop_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := FromSlice([]int{1, 2, 3, 4, 5}).Drop(2)
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	assert.Equal(t, []int{0, 1, 2}, indexes, "Expected %v, got %v", []int{0, 1, 2}, indexes)
}
