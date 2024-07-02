package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	t.Parallel()
	want := []int{}
	slice := Empty[int]().Collect()
	assert.Equal(t, want, slice, "Expected %v, got %v", want, slice)
}
