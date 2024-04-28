package chapter4

import "math"

func Factorial(number int) int {
	if number <= 0 {
		return 1
	}

	return number * Factorial(number-1)
}

func ConstantE(number int) float64 {
	if number <= 0 {
		return 1.0
	}
	e := 1.0

	for counter := 1; counter <= number; counter++ {
		e += 1.0 / float64(Factorial(counter))
	}
	return e
}

func ConstantEx(number int) float64 {
	if number <= 0 {
		return 1.0
	}
	return math.Pow(ConstantE(number), float64(number))
}
