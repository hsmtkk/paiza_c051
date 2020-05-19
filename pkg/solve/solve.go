package solve

import (
	"github.com/hsmtkk/paiza_c051/pkg/permutation"
	"github.com/hsmtkk/paiza_c051/pkg/score"
)

func Solve(numbers []int) int {
	perms := permutation.Permutation(numbers)
	max := score.MaxScore(perms)
	return max
}
