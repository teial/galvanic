package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice Slice[int]
		other Slice[int]
		want  []Pair[int, int]
	}{
		{
			name:  "empty",
			slice: []int{},
			other: []int{},
			want:  []Pair[int, int]{},
		},
		{
			name:  "single",
			slice: []int{1},
			other: []int{2},
			want:  []Pair[int, int]{NewPair(1, 2)},
		},
		{
			name:  "multiple",
			slice: []int{1, 2, 3},
			other: []int{4, 5, 6},
			want:  []Pair[int, int]{NewPair(1, 4), NewPair(2, 5), NewPair(3, 6)},
		},
		{
			name:  "different length",
			slice: []int{1, 2, 3},
			other: []int{4, 5},
			want:  []Pair[int, int]{NewPair(1, 4), NewPair(2, 5)},
		},
		{
			name:  "different length 2",
			slice: []int{1, 2},
			other: []int{4, 5, 6},
			want:  []Pair[int, int]{NewPair(1, 4), NewPair(2, 5)},
		},
		{
			name:  "different length 3",
			slice: []int{1, 2, 3},
			other: []int{},
			want:  []Pair[int, int]{},
		},
		{
			name:  "different length 4",
			slice: []int{},
			other: []int{4, 5, 6},
			want:  []Pair[int, int]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Zip(tc.slice, tc.other).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestZip_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Zip(Slice[int]{1, 2, 3}, Slice[int]{1, 2})
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	assert.Equal(t, []int{0, 1}, indexes, "Expected %v, got %v", []int{0, 1}, indexes)
}
