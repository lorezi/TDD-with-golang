package arraysslices

// Sum will take an array of numbers and return the total
func Sum(xs [5]int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func SumN(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}
