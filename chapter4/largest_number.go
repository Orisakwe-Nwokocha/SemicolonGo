package chapter4

import (
	"fmt"
	"math"
	"semicolonGo/utils"
)

func LargestNumber() {
	input := utils.Input{}

	largest := math.MinInt32
	counter := 0
	var number int

	for counter < 3 {
		fmt.Println("Enter number: ")
		num, successful := input.NextInt()
		if !successful {
			fmt.Println("Invalid input")
			continue
		}

		if num > largest {
			largest = num
		}
		counter++
		number = num
		fmt.Println("counter:", counter)
		fmt.Println("num:", number)
	}

	fmt.Printf("\n\nMost recent number: %d\n", number)
	fmt.Printf("The largest number: %d\n", largest)
	fmt.Printf("Number of times inputed: %d\n", counter)
}
