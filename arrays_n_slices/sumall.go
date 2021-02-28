package arraysandslices

// SumAll the slices parameter into a single slice containing sum of each of the slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		sums = append(sums, sum)
	}
	return
}
