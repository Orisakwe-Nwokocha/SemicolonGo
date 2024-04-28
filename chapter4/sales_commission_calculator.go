package chapter4

import (
	"fmt"
	"semicolonGo/utils"
)

func SalesCommission() {
	input := utils.Input{}

	wage := 200.0
	item1 := 239.99
	item2 := 129.75
	item3 := 99.95
	item4 := 350.89

	item1Count := 0
	item2Count := 0
	item3Count := 0
	item4Count := 0

	for {
		fmt.Println("Enter item ID sold (1-4) or 0 to end: ")
		number, _ := input.NextInt()
		if number == 0 {
			break
		}

		switch number {
		case 1:
			item1Count++
		case 2:
			item2Count++
		case 3:
			item3Count++
		case 4:
			item4Count++
		}
	}

	result1 := float64(item1Count) * item1
	result2 := float64(item2Count) * item2
	result3 := float64(item3Count) * item3
	result4 := float64(item4Count) * item4
	resultTotal := result1 + result2 + result3 + result4

	fullWage := wage + (0.09 * resultTotal)

	fmt.Printf("\n\nThe salesperson's commission for the week is $%.2f\n", fullWage)
}
