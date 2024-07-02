package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Iterable[int]
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
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(int) bool { return false },
			want: []int{},
		},
		{
			name: "some true",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i < 3 },
			want: []int{1, 2},
		},
		{
			name: "some false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i > 3 },
			want: []int{},
		},
		{
			name: "mixed",
			seq:  Values(2, 4, 1, 3, 6, 5),
			fn:   func(i int) bool { return i%2 == 0 },
			want: []int{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.All().TakeWhile(tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestTakeWhile_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  ReverseIterable[int]
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
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "all false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(int) bool { return false },
			want: []int{},
		},
		{
			name: "some true",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i > 3 },
			want: []int{5, 4},
		},
		{
			name: "some false",
			seq:  Values(1, 2, 3, 4, 5),
			fn:   func(i int) bool { return i < 3 },
			want: []int{},
		},
		{
			name: "mixed",
			seq:  Values(1, 3, 6, 5, 2, 4),
			fn:   func(i int) bool { return i%2 == 0 },
			want: []int{4, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.Reverse().TakeWhile(tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestTakeWhile_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Values(1, 3, 5, 2, 4).Reverse().TakeWhile(func(e int) bool { return e%2 == 0 })
	for i, e := range sequence.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
