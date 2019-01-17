package reverse_string

import (
	"fmt"
	"testing"
)

func TestReverseBasic(t *testing.T) {
	fmt.Printf("input = %s, output = %s", "cool", reverse("cool"))
}

func TestReverseTable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"cool", "looc"},
		{"internal test", "tset lanretni"},
		{"someh i i i ", " i i i hemos"},
		{"aha i9p qqa", "aqq p9i aha"},
	}
	for _, test := range tests {
		if output := reverse(test.input); output != test.expected {
			t.Errorf("Error, expected: '%s' but got '%s'", test.expected, output)
		}
	}
}
