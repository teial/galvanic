package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		slice Slice[string]
		fn    func(string) int
		want  []int
	}{
		{
			name:  "empty",
			slice: Slice[string]{},
			fn:    func(string) int { return 0 },
			want:  []int{},
		},
		{
			name:  "single",
			slice: Slice[string]{"a"},
			fn:    func(string) int { return 1 },
			want:  []int{1},
		},
		{
			name:  "multiple",
			slice: Slice[string]{"a", "ab", "abc"},
			fn:    func(x string) int { return len(x) },
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Map(tc.slice, tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}
