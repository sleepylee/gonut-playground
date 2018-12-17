package permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationFullCorrect(t *testing.T) {
	count := 0
	Perm([]rune("abcd"), func(a []rune) {
		//fmt.Println(string(a))
		count++
	})
	assert.Equal(t, count, 24, "amount of full pair permutations P(4,4)")
}
