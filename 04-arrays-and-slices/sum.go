package arraysandslices

func Sum(numbers []int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}
