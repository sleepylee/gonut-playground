package perfect_number

import (
	"errors"
	"fmt"
)

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationDeficient
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Can't classify zero or negative.")

// Use unit64 because we only validate positive ( > 0) input
func Classify(input uint64) (Classification, error) {
	if input <= 0 {
		return 0, ErrOnlyPositive
	}

	aliquotSum := aliquotSumOfDivisor(input)
	if aliquotSum == input {
		return ClassificationPerfect, nil
	} else if aliquotSum < input {
		return ClassificationDeficient, nil
	} else {
		return ClassificationAbundant, nil
	}
}

func aliquotSumOfDivisor(inp uint64) uint64 {
	var result uint64
	for _, num := range findAllPositiveDivisor(inp) {
		result += num
	}
	return result
}

func findAllPositiveDivisor(self uint64) []uint64 {
	var ui64 uint64 = 1
	var result []uint64

	for ui64 <= self/2 {
		if self%ui64 == 0 {
			result = append(result, ui64)
		}
		ui64 += 1
	}
	fmt.Println(result)
	return result
}
