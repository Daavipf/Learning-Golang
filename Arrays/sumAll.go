package arrays

func SumAll(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
