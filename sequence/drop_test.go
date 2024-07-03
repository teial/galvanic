package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Iterable[int]
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
			want: []int{},
		},
		{
			name: "single 0",
			seq:  Values(1),
			n:    0,
			want: []int{1},
		},
		{
			name: "single 2",
			seq:  Values(1),
			n:    2,
			want: []int{},
		},
		{
			name: "multiple",
			seq:  Values(1, 2, 3),
			n:    2,
			want: []int{3},
		},
		{
			name: "multiple 0",
			seq:  Values(1, 2, 3),
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "multiple 4",
			seq:  Values(1, 2, 3),
			n:    4,
			want: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.All().Drop(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestDrop_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  ReverseIterable[int]
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
			want: []int{},
		},
		{
			name: "single 0",
			seq:  Values(1),
			n:    0,
			want: []int{1},
		},
		{
			name: "single 2",
			seq:  Values(1),
			n:    2,
			want: []int{},
		},
		{
			name: "multiple",
			seq:  Values(1, 2, 3),
			n:    2,
			want: []int{1},
		},
		{
			name: "multiple 0",
			seq:  Values(1, 2, 3),
			n:    0,
			want: []int{3, 2, 1},
		},
		{
			name: "multiple 4",
			seq:  Values(1, 2, 3),
			n:    4,
			want: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.Reverse().Drop(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestDrop_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range Values(1, 2, 3, 4, 5).Drop(2).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}

func TestDrop_Indexes_Reverse(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range Values(1, 2, 3, 4, 5).Reverse().Drop(2).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
