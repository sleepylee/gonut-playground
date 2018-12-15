package main

import "fmt"

func main() {
	permuCount := 0
	Perm([]rune("abcd"), func(a []rune) {
		fmt.Println(string(a))
		permuCount++
	})
	fmt.Printf("Done! P(n, r) = %d \n", permuCount)

	permuRepCount := 0
	for combination := range genPermutationRep("abcd", 3) {
		fmt.Println(combination)
		permuRepCount++
	}
	fmt.Printf("Done! P-rep(n, r) = %d", permuRepCount)
}
