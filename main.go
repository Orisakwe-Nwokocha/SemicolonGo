package main

import (
	"fmt"
	"time"
)

func main() {

	//chapter4.CalculateGasMileage()
	//chapter4.CreditLimit()
	//chapter4.SalesCommission()
	//chapter4.LargestNumber()
	//fmt.Println(chapter4.Factorial(-5))
	//fmt.Println(chapter4.Factorial(5))
	//fmt.Println(chapter4.ConstantE(3))
	//fmt.Println(chapter4.ConstantEx(3))
	//fmt.Println(chapter4.ConstantE(5))
	//fmt.Println(chapter4.ConstantEx(5))
	//chapter4.WorldPopulationGrowthDifference()
	//chapter4.Encrypt()
	//chapter4.Decrypt()
	//chapter5.PrintTriangles()
	//fmt.Println(sort([]int{5, 1, 3, 2, 4}))
	//chapter5.PrintDiamond()
	//chapter5.CalculateFairTax()
	//chapter5.TakeGlobalWarmingQuiz()
	//checkoutapp.Checkout()

	fmt.Println(time.Now().Format("02-Jan-06 3:04:05 pm"))
}

func sort(numbers []int) []int {
	for index, number := range numbers {
		numbers[index] = number * number

	}
	for counter := 0; counter < len(numbers); counter++ {
		for index := 0; index < len(numbers)-1; index++ {
			if numbers[index] > numbers[index+1] {
				numbers[index], numbers[index+1] = numbers[index+1], numbers[index]
			}
		}

	}
	return numbers
}
