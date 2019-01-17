package atbash_cipher

import (
	"fmt"
	"testing"
)

func TestCipherGivesCorrectAnswers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "z"},
		{"test", "gvhg"},
		{"testwithspace", "gvhgd rgshk zxv"},
	}
	for _, test := range tests {
		if cip := cipherChar(test.input); cip != test.expected {
			t.Errorf("Error on cipher, expected: '%s' but got '%s'", test.expected, cip)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "z"},
		{"gvhg", "test"},
		{"gvhgd rgshk zxv", "testwithspace"},
		{"gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt", "thequickbrownfoxjumpsoverthelazydog"},
	}
	for _, test := range tests {
		if dec := decodeChar(test.input); dec != test.expected {
			t.Errorf("Error on decode, expected: '%s' but got '%s'", test.expected, dec)
		}
	}
}

func TestRun(t *testing.T) {
	input := "thequickbrownfoxjumpsoverthelazydog"
	cipher := cipherChar(input)
	decode := decodeChar(cipher)
	fmt.Printf("input = %s\ncipher = %s\ncecode = %s\n", input, cipher, decode)
	if input != decode {
		t.Error("Wronggg")
	}

}
