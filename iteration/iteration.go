package iteration

// Iteration takes a character and number of iteration and returns (n)times of the character
func Iteration(character string, n int) string {
	repeated := ""
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
