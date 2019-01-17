package reverse_string

func reverse(input string) string {
	length := len(input)
	output := make([]rune, length)
	for i, c := range input {
		output[length-i-1] = c
	}
	return string(output)
}
