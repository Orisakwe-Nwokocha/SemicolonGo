package mylambdas

func Map(numbers []int, function func(int) int) []int {
	result := make([]int, len(numbers))
	for i, number := range numbers {
		result[i] = function(number)
	}
	return result
}
