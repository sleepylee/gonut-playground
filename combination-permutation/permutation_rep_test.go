package permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationCorrect(t *testing.T) {
	count := 0
	for _ = range genPermutationRep("abcd", 3) {
		//fmt.Println(v)
		count++
	}
	assert.Equal(t, count, 64, "amount of permutations P(4,3)")
}
