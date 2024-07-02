package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Sequence[int]
		want []int
	}{
		{
			name: "empty",
			seq:  Empty[int](),
			want: []int{},
		},
		{
			name: "single",
			seq:  Values(1),
			want: []int{1},
		},
		{
			name: "multiple",
			seq:  Values(1, 2, 3),
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

func TestValues_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range Values(1, 2, 3).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
