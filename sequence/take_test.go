package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	tests := []struct {
		name  string
		slice Slice[int]
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
			want:  []int{1},
		},
		{
			name:  "single 0",
			slice: []int{1},
			n:     0,
			want:  []int{},
		},
		{
			name:  "single 2",
			slice: []int{1},
			n:     2,
			want:  []int{1},
		},
		{
			name:  "multiple",
			slice: []int{1, 2, 3},
			n:     2,
			want:  []int{1, 2},
		},
		{
			name:  "multiple 0",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{},
		},
		{
			name:  "multiple 4",
			slice: []int{1, 2, 3},
			n:     4,
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.slice.All().Take(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestTake_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Slice[int]{1, 2, 3, 4, 5}.All().Take(3)
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	assert.Equal(t, []int{0, 1, 2}, indexes, "Expected %v, got %v", []int{0, 1, 2}, indexes)
}
