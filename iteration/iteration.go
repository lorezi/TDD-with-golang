package iteration

// Iteration takes a character and returns x5 of the character
func Iteration(character string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
