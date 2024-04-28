package mylambdas

func Filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, number := range numbers {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return result
}
