package hamming_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingCountCorrect(t *testing.T) {
	result, _ := hammingCount("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	assert.Equal(t, 7, result, "Amount of different pair")

	result2, _ := hammingCount("GAGCCTACTAACGGGAT", "GAGCCTACTAACGGGAT")
	assert.Equal(t, 0, result2, "There is no difference")
}

func TestHammingCountThrowError(t *testing.T) {
	_, error := hammingCount("GAGCCTACTAACGGGAT", "CATCGTAATGACGCT")
	assert.IsType(t, NotMatchError{}, error, "Error is a NotMatchError")
	assert.Equal(t, "length not match: 0 vs 0", error.Error(), "Error message is matched as designed")
}
