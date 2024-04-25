package arraysandslices

func Sum(numbers [5]int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}
