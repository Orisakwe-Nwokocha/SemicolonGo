package chapter5

import "fmt"

func PrintTriangles() {
	for counter := 1; counter <= 10; counter++ {
		printStars(counter)
		printSpaces(14 - counter)
		printStars(11 - counter)
		printSpaces((counter + 1) * 2)
		printStars(11 - counter)
		printSpaces(14 - counter)
		printStars(counter)
		fmt.Println()
	}
}

func PrintDiamond() {
	stars := 1
	spaces := 4
	for counter := 1; counter <= 9; counter++ {
		printSpaces(spaces)
		printStars(stars)
		fmt.Println()
		if counter >= 5 {
			stars -= 2
			spaces++
			continue
		}
		stars += 2
		spaces--
	}
}

func printSpaces(number int) {
	for row := 1; row <= number; row++ {
		fmt.Print(" ")
	}
}

func printStars(number int) {
	for row := 1; row <= number; row++ {
		fmt.Print("*")
	}
}
