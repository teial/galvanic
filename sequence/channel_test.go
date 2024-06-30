package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeChannel(nums ...int) Channel[int] {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			ch <- num
		}
		close(ch)
	}()
	return Channel[int](ch)
}

func TestChannel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		ch   Channel[int]
		want []int
	}{
		{
			name: "empty",
			ch:   makeChannel(),
			want: []int{},
		},
		{
			name: "single",
			ch:   makeChannel(1),
			want: []int{1},
		},
		{
			name: "multiple",
			ch:   makeChannel(1, 2, 3),
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.ch.All().Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}
