package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from int
		to   int
		want []int
	}{
		{
			name: "empty",
			from: 0,
			to:   0,
			want: []int{},
		},
		{
			name: "single",
			from: 1,
			to:   2,
			want: []int{1},
		},
		{
			name: "multiple",
			from: 1,
			to:   4,
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := RangeTo(tc.from, tc.to).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestRangeTo_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from int
		to   int
		want []int
	}{
		{
			name: "empty",
			from: 0,
			to:   0,
			want: []int{},
		},
		{
			name: "single",
			from: 1,
			to:   2,
			want: []int{1},
		},
		{
			name: "multiple",
			from: 1,
			to:   4,
			want: []int{3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := RangeTo(tc.from, tc.to).Reverse().Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestRangeTo_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range RangeTo(0, 3).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}

func TestRangeTo_Indexes_Reverse(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range RangeTo(0, 3).Reverse().Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}

func TestRange(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from int
		n    int
		want []int
	}{
		{
			name: "empty",
			from: 0,
			n:    0,
			want: []int{},
		},
		{
			name: "single",
			from: 1,
			n:    3,
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Range(tc.from).Take(tc.n).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestRange_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range Range(0).Take(3).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
