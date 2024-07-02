package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "empty",
			slice: []int{},
			want:  []int{},
		},
		{
			name:  "single",
			slice: []int{1},
			want:  []int{1},
		},
		{
			name:  "multiple",
			slice: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := FromSlice(tc.slice).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestSlice_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	for i, e := range FromSlice([]int{1, 2, 3}).Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
