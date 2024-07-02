package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		seq  Iterable[string]
		fn   func(string) int
		want []int
	}{
		{
			name: "empty",
			seq:  Empty[string](),
			fn:   func(string) int { return 0 },
			want: []int{},
		},
		{
			name: "single",
			seq:  Values("a"),
			fn:   func(string) int { return 1 },
			want: []int{1},
		},
		{
			name: "multiple",
			seq:  Values("a", "ab", "abc"),
			fn:   func(s string) int { return len(s) },
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Map(tc.seq, tc.fn).Collect()
			assert.Equal(t, tc.want, slice, "Expected %v, got %v", tc.want, slice)
		})
	}
}

func TestMap_Indexes(t *testing.T) {
	t.Parallel()
	indexes := make([]int, 0)
	sequence := Values("a", "ab", "abc")
	mapped := Map(sequence, func(e string) int { return len(e) })
	for i, e := range mapped.Fn2 {
		_ = e
		indexes = append(indexes, i)
	}
	want := []int{0, 1, 2}
	assert.Equal(t, want, indexes, "Expected %v, got %v", want, indexes)
}
