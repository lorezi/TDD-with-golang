package arraysslices

// Sum will take an array of numbers and return the total
func Sum(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

// SumAll will take a varying number of slices (variadic function) and return a new slice containing the total for each slice passed in
func SumAll(xs ...[]int) []int {

	xsum := []int{}

	for _, v := range xs {
		total := 0
		for _, n := range v {
			total += n

		}
		xsum = append(xsum, total)
	}

	return xsum
}
