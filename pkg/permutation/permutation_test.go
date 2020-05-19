package permutation_test

import (
	"testing"

	"github.com/hsmtkk/paiza_c051/pkg/permutation"
	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	want := [][]int{}
	want = append(want, []int{1, 2, 3})
	want = append(want, []int{1, 3, 2})
	want = append(want, []int{2, 1, 3})
	want = append(want, []int{2, 3, 1})
	want = append(want, []int{3, 1, 2})
	want = append(want, []int{3, 2, 1})
	got := permutation.Permutation([]int{1, 2, 3})
	assert.Equal(t, len(want), len(got), "should be equal")
	for _, g := range got {
		assert.Contains(t, want, g, "should contains")
	}
}
