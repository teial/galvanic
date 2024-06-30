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
