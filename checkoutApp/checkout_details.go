package checkoutapp

import (
	"semicolonGo/utils"
	"strings"
)

func collectCheckoutDetails() ([]string, []int, []float64) {
	var items []string
	var purchasedQuantities []int
	var itemPrices []float64

	for {
		item := input("What the user buy?")
		items = append(items, item)
		purchasedQuantities = append(purchasedQuantities, getQuantityPurchased())
		itemPrices = append(itemPrices, getPriceOfItem())

		sentinel := getSentinelValue()
		if sentinel == "no" {
			break
		}
	}

	return items, purchasedQuantities, itemPrices
}

func getSentinelValue() string {
	for {
		value := input("Add more items (yes/no)?")
		value = strings.ToLower(value)

		if value != "yes" && value != "no" {
			display("Invalid input")
			continue
		}

		return value
	}

}

func getPriceOfItem() float64 {
	scanner := utils.Input{}

	for {
		display("How much per unit?")
		price, isSuccessful := scanner.NextFloat()
		if !isSuccessful {
			display("Invalid input")
			continue
		}

		return price
	}

}

func getQuantityPurchased() int {
	scanner := utils.Input{}

	for {
		display("How many pieces?")
		quantity, isSuccessful := scanner.NextInt()
		if !isSuccessful {
			display("Invalid input")
			continue
		}

		return quantity
	}

}

func getDiscount() float64 {
	scanner := utils.Input{}

	for {
		display("How much discount will the customer get?")
		discount, isSuccessful := scanner.NextFloat()

		isInvalid := discount < 0 || discount > 100
		if !isSuccessful || isInvalid {
			display("Invalid input")
			continue
		}

		return discount
	}
}
