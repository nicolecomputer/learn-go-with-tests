package arraysandslices

func Sum(numbers []int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}
