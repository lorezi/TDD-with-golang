package iteration

// Iteration takes a character and returns x5 of the character
func Iteration(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
