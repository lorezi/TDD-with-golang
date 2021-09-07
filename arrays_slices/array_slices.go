package arraysslices

// Sum will take an array of numbers and return the total
func Sum(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}
