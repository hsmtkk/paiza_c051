package solve_test

import (
	"testing"

	"github.com/hsmtkk/paiza_c051/pkg/solve"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	want := 175
	got := solve.Solve([]int{9, 2, 3, 8})
	assert.Equal(t, want, got, "should be equal")
}
