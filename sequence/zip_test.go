package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq1 Sequence[int]
		seq2 Sequence[int]
		want []Pair[int, int]
	}{
		{
			name: "empty",
			seq1: Empty[int](),
			seq2: Empty[int](),
			want: []Pair[int, int]{},
		},
		{
			name: "single",
			seq1: Values(1),
			seq2: Values(2),
			want: []Pair[int, int]{NewPair(1, 2)},
		},
		{
			name: "multiple",
			seq1: Values(1, 2, 3),
			seq2: Values(4, 5, 6),
			want: []Pair[int, int]{NewPair(1, 4), NewPair(2, 5), NewPair(3, 6)},
		},
		{
			name: "different length",
			seq1: Values(1, 2, 3),
			seq2: Values(4, 5),
			want: []Pair[int, int]{NewPair(1, 4), NewPair(2, 5)},
		},
		{
			name: "different length 2",
			seq1: Values(1, 2),
			seq2: Values(4, 5, 6),
			want: []Pair[int, int]{NewPair(1, 4), NewPair(2, 5)},
		},
		{
			name: "different length 3",
			seq1: Values(1, 2, 3),
			seq2: Empty[int](),
			want: []Pair[int, int]{},
		},
		{
			name: "different length 4",
			seq1: Empty[int](),
			seq2: Values(4, 5, 6),
			want: []Pair[int, int]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			zipped := Zip(tc.seq1, tc.seq2).Collect()
			assert.Equal(t, tc.want, zipped, "Expected %v, got %v", tc.want, zipped)
		})
	}
}

func TestZip_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	slice := Values(1, 2, 3)
	other := Values(1, 2)
	for i, e := range Zip(slice, other).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
