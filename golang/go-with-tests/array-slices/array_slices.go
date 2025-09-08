package arrayslices


func Sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
