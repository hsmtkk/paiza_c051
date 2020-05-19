package score_test

import(
	"testing"
	"github.com/hsmtkk/paiza_c051/pkg/score"
	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T){
	want := 175
	got := score.Score([]int{9,2,8,3})
	assert.Equal(t, want, got, "should be equal")
}