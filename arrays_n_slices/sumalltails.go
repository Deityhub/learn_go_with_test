package arraysandslices

// SumAllTails calculates the totals of the "tails" of each slice
func SumAllTails(numbersToTail ...[]int) (tails []int) {
	for _, numbers := range numbersToTail {
		if len(numbers) == 0 {
			numbers = []int{0}
		} else {
			numbers = numbers[1:]
		}
		tails = append(tails, Sum(numbers))
	}
	return
}
