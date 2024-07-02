package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeChannel(nums ...int) Sequence[int] {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			ch <- num
		}
		close(ch)
	}()
	return Channel(ch)
}

func TestChannel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Sequence[int]
		want []int
	}{
		{
			name: "empty",
			seq:  makeChannel(),
			want: []int{},
		},
		{
			name: "single",
			seq:  makeChannel(1),
			want: []int{1},
		},
		{
			name: "multiple",
			seq:  makeChannel(1, 2, 3),
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.seq.Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestChannel_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range makeChannel(1, 2, 3).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
