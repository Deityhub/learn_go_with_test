package iteration

// Iterate returns word any number of times specified
func Iterate(word string, iteration int) string {
	if iteration <= 0 {
		iteration = 1
	}

	var output string
	for i := 0; i < iteration; i++ {
		output += word
	}
	return output
}
