package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipAll(t *testing.T) {
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
			want: []Pair[int, int]{NewPair(1, 4), NewPair(2, 5), {3, true, 0, false}},
		},
		{
			name: "different length 2",
			seq1: Values(1, 2),
			seq2: Values(4, 5, 6),
			want: []Pair[int, int]{NewPair(1, 4), NewPair(2, 5), {0, false, 6, true}},
		},
		{
			name: "different length 3",
			seq1: Values(1, 2),
			seq2: Empty[int](),
			want: []Pair[int, int]{{1, true, 0, false}, {2, true, 0, false}},
		},
		{
			name: "different length 4",
			seq1: Empty[int](),
			seq2: Values(4, 5),
			want: []Pair[int, int]{{0, false, 4, true}, {0, false, 5, true}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			zipped := ZipAll(tc.seq1, tc.seq2).Collect()
			assert.Equal(t, tc.want, zipped, "Expected %v, got %v", tc.want, zipped)
		})
	}
}

func TestZipAll_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	seq1 := Values(1, 2, 3)
	seq2 := Values(1, 2)
	for i, e := range ZipAll(seq1, seq2).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
