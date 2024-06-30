package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice Slice[int]
		fn    func(int) bool
		want  []int
	}{
		{
			name:  "empty",
			slice: []int{},
			fn:    func(int) bool { return true },
			want:  []int{},
		},
		{
			name:  "single",
			slice: []int{1},
			fn:    func(int) bool { return true },
			want:  []int{1},
		},
		{
			name:  "single false",
			slice: []int{1},
			fn:    func(int) bool { return false },
			want:  []int{},
		},
		{
			name:  "multiple",
			slice: []int{1, 2, 3},
			fn:    func(int) bool { return true },
			want:  []int{1, 2, 3},
		},
		{
			name:  "multiple false",
			slice: []int{1, 2, 3},
			fn:    func(int) bool { return false },
			want:  []int{},
		},
		{
			name:  "multiple odd",
			slice: []int{1, 2, 3},
			fn:    func(e int) bool { return e%2 == 1 },
			want:  []int{1, 3},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := tc.slice.All().Filter(tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}
