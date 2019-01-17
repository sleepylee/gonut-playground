package perfect_number

import (
	"testing"
)

var classificationTestCases = []struct {
	input    uint64
	expected Classification
}{
	{1, ClassificationDeficient},
	{2, ClassificationDeficient},
	{4, ClassificationDeficient},
	{13, ClassificationDeficient},
	{12, ClassificationAbundant},
	{18, ClassificationAbundant},
	{20, ClassificationAbundant},
	{6, ClassificationPerfect},
	{28, ClassificationPerfect},
	{496, ClassificationPerfect},
	{8128, ClassificationPerfect},
}

func TestGivesPositiveRequiredError(t *testing.T) {
	if _, err := Classify(0); err != ErrOnlyPositive {
		t.Errorf("expected error %q but got %q", ErrOnlyPositive, err)
	}
}

func TestClassifiesCorrectly(t *testing.T) {
	for _, c := range classificationTestCases {
		if cat, err := Classify(c.input); err != nil {
			t.Errorf("at input = %d: Expected no error but got %s", c.input, err)
		} else if cat != c.expected {
			t.Errorf("at input = %d: Expected %q, got %q", c.input, c.expected, cat)
		}
	}
}
