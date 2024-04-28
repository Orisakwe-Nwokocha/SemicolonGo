package chapter4

import (
	"fmt"
	"semicolonGo/utils"
)

func CalculateGasMileage() {
	input := utils.Input{}

	totalMiles := 0
	totalGallons := 0
	totalTrips := 0
	combined := 0.0

	for {
		fmt.Println("Enter miles driven or enter -1 to quit: ")
		miles, _ := input.NextInt()
		if miles == -1 {
			break
		}

		fmt.Println("Enter gallons used: ")
		gallons, _ := input.NextInt()

		milesPerGallon := float64(miles) / float64(gallons)
		fmt.Printf("The miles per gallon for this trip: %.2fMPG\n", milesPerGallon)

		combined += milesPerGallon
		totalMiles += miles
		totalGallons += gallons
		totalTrips++
	}

	combinedMpg := float64(totalMiles) / float64(totalGallons)

	format := "\n\nThe total miles driven for %d trips: %d miles\nThe total gallons used for %d trips: %d gallons\n"
	fmt.Printf(format, totalTrips, totalMiles, totalTrips, totalGallons)
	fmt.Printf("\nThe combined miles per gallon obtained for %d trips: %.2fMPG\n", totalTrips, combinedMpg)
	fmt.Printf("Type 2: The combined miles per gallon obtained for %d trips: %.2fMPG\n", totalTrips, combined)
}
