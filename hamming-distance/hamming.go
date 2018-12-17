package main

import "fmt"

type NotMatchError struct {
	requiredLength int
	inputLength    int
}

func (nme NotMatchError) Error() string {
	return fmt.Sprintf("length not match: %d vs %d", nme.requiredLength, nme.inputLength)
}

func hammingCount(spec string, dna string) (int, error) {
	if len(spec) != len(dna) {
		return len(dna), NotMatchError{}
	}
	count := 0
	for i, each := range spec {
		if each != []rune(dna)[i] {
			count++
		}
	}
	return count, nil
}

func main() {
	_, err := hammingCount("GAGCCTACTAACGGG", "CATCGTAATGACGGCCT")
	if err != nil {
		fmt.Println(err)
	}

	result, err := hammingCount("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("diff = %d ", result)
}
