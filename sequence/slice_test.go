package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice Slice[int]
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
			slice := tc.slice.All().Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}
