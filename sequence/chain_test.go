package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		first Iterable[int]
		seqs  []Iterable[int]
		want  []int
	}{
		{
			name:  "first empty seqs empty",
			first: Empty[int](),
			seqs:  []Iterable[int]{},
			want:  []int{},
		},
		{
			name:  "first empty seqs single",
			first: Empty[int](),
			seqs:  []Iterable[int]{Values(1)},
			want:  []int{1},
		},
		{
			name:  "first empty seqs multiple",
			first: Empty[int](),
			seqs:  []Iterable[int]{Values(1, 2, 3)},
			want:  []int{1, 2, 3},
		},
		{
			name:  "first empty seqs mixed",
			first: Empty[int](),
			seqs: []Iterable[int]{
				Empty[int](),
				Values(1),
				Empty[int](),
				Values(2, 3),
				Empty[int](),
			},
			want: []int{1, 2, 3},
		},
		{
			name:  "first single seqs empty",
			first: Values(1),
			seqs:  []Iterable[int]{},
			want:  []int{1},
		},
		{
			name:  "first single seqs single",
			first: Values(1),
			seqs:  []Iterable[int]{Values(2)},
			want:  []int{1, 2},
		},
		{
			name:  "first single seqs multiple",
			first: Values(1),
			seqs:  []Iterable[int]{Values(2, 3, 4)},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "first single seqs mixed",
			first: Values(1),
			seqs: []Iterable[int]{
				Empty[int](),
				Values(2),
				Empty[int](),
				Values(3, 4),
				Empty[int](),
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name:  "first multiple seqs empty",
			first: Values(1, 2, 3),
			seqs:  []Iterable[int]{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "first multiple seqs single",
			first: Values(1, 2, 3),
			seqs:  []Iterable[int]{Values(4)},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "first multiple seqs multiple",
			first: Values(1, 2, 3),
			seqs:  []Iterable[int]{Values(4, 5, 6)},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "first multiple seqs mixed",
			first: Values(1, 2, 3),
			seqs: []Iterable[int]{
				Empty[int](),
				Values(4),
				Empty[int](),
				Values(5, 6),
				Empty[int](),
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.first.All().Chain(tc.seqs...).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestChain_Reverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		first ReverseIterable[int]
		seqs  []ReverseIterable[int]
		want  []int
	}{
		{
			name:  "first empty seqs empty",
			first: Empty[int](),
			seqs:  []ReverseIterable[int]{},
			want:  []int{},
		},
		{
			name:  "first empty seqs single",
			first: Empty[int](),
			seqs:  []ReverseIterable[int]{Values(1)},
			want:  []int{1},
		},
		{
			name:  "first empty seqs multiple",
			first: Empty[int](),
			seqs:  []ReverseIterable[int]{Values(1, 2, 3)},
			want:  []int{3, 2, 1},
		},
		{
			name:  "first empty seqs mixed",
			first: Empty[int](),
			seqs: []ReverseIterable[int]{
				Empty[int](),
				Values(1),
				Empty[int](),
				Values(2, 3),
				Empty[int](),
			},
			want: []int{3, 2, 1},
		},
		{
			name:  "first single seqs empty",
			first: Values(1),
			seqs:  []ReverseIterable[int]{},
			want:  []int{1},
		},
		{
			name:  "first single seqs single",
			first: Values(1),
			seqs:  []ReverseIterable[int]{Values(2)},
			want:  []int{2, 1},
		},
		{
			name:  "first single seqs multiple",
			first: Values(1),
			seqs:  []ReverseIterable[int]{Values(2, 3, 4)},
			want:  []int{4, 3, 2, 1},
		},
		{
			name:  "first single seqs mixed",
			first: Values(1),
			seqs: []ReverseIterable[int]{
				Empty[int](),
				Values(2),
				Empty[int](),
				Values(3, 4),
				Empty[int](),
			},
			want: []int{4, 3, 2, 1},
		},
		{
			name:  "first multiple seqs empty",
			first: Values(1, 2, 3),
			seqs:  []ReverseIterable[int]{},
			want:  []int{3, 2, 1},
		},
		{
			name:  "first multiple seqs single",
			first: Values(1, 2, 3),
			seqs:  []ReverseIterable[int]{Values(4)},
			want:  []int{4, 3, 2, 1},
		},
		{
			name:  "first multiple seqs multiple",
			first: Values(1, 2, 3),
			seqs:  []ReverseIterable[int]{Values(4, 5, 6)},
			want:  []int{6, 5, 4, 3, 2, 1},
		},
		{
			name:  "first multiple seqs mixed",
			first: Values(1, 2, 3),
			seqs: []ReverseIterable[int]{
				Empty[int](),
				Values(4),
				Empty[int](),
				Values(5, 6),
				Empty[int](),
			},
			want: []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.first.AllR().Chain(tc.seqs...).Reverse().Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestChain_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Values(
		1,
		2,
	).Chain(Empty[int](), Values(3), Empty[int](), Values(4, 5), Empty[int]())
	for i, e := range sequence.All().Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2, 3, 4}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
