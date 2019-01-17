package atbash_cipher

// TODO: cipher & decode with number, capital char as well yeah?
func cipherChar(input string) string {
	var output string

	for i, c := range input {
		if i != 0 && i%5 == 0 {
			output += " "
		}
		revPos := 219 - c // 97 = a, 122 = z
		output = output + string(revPos)
	}
	return output
}

func decodeChar(input string) string {
	var output string
	for _, c := range input {
		if c != 32 {
			revPos := 219 - c // 97 = a, 122 = z
			output = output + string(revPos)
		}
	}
	return output
}
