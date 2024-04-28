package chapter4

import (
	"fmt"
	"regexp"
	"semicolonGo/utils"
)

func WorldPopulationGrowthDifference() {
	const InitialPopulation = 8072783338
	oldPopulation := InitialPopulation
	const GrowthRate = 0.88 / 100
	doubleYear := 0
	yearCounter := 1

	fmt.Println("Year\t   Anticipated Population\t\tNumerical Increase")
	for yearCounter <= 75 {
		newPopulation := int(float64(oldPopulation)*GrowthRate) + oldPopulation
		growthDifference := newPopulation - oldPopulation
		oldPopulation = newPopulation
		fmt.Printf("%d\t%23d\t   %24d\n", yearCounter, newPopulation, growthDifference)

		yearCounter++

		isDouble := newPopulation >= (InitialPopulation * 2)
		if isDouble && doubleYear == 0 {
			doubleYear = yearCounter
		}
	}

	if doubleYear == 0 {
		fmt.Printf("\nPopulation would not be doubled within %d years\n", yearCounter-1)
	} else {
		fmt.Printf("\nThe population would be doubled in Year %d\n", doubleYear)
	}
}

func Encrypt() {
	input := utils.Input{}
	var number string
	fmt.Printf("Enter a 4 digit number: ")
	number = input.NextLine()

	for length := len(number); length != 4 || !regexp.MustCompile(`^\d{4}$`).MatchString(number); {
		fmt.Printf("%s is not a valid 4 digit number\n", number)
		fmt.Printf("Enter a 4 digit number: ")
		number = input.NextLine()
		length = len(number)
	}

	a := number[0]
	b := number[1]
	c := number[2]
	d := number[3]

	num1 := int(a - '0')
	num2 := int(b - '0')
	num3 := int(c - '0')
	num4 := int(d - '0')

	num1 = (num1 + 7) % 10
	num2 = (num2 + 7) % 10
	num3 = (num3 + 7) % 10
	num4 = (num4 + 7) % 10

	encryptedValue := fmt.Sprintf("%d%d%d%d", num3, num4, num1, num2)

	fmt.Printf("\nThe encrypted value of %s is %s\n", number, encryptedValue)
}

func Decrypt() {
	input := utils.Input{}
	var number string
	fmt.Printf("Enter a 4 digit number: ")
	number = input.NextLine()

	for length := len(number); length != 4 || !regexp.MustCompile(`^\d{4}$`).MatchString(number); {
		fmt.Printf("%s is not a valid 4 digit number\n", number)
		fmt.Printf("Enter a 4 digit number: ")
		number = input.NextLine()
		length = len(number)
	}

	a := number[0]
	b := number[1]
	c := number[2]
	d := number[3]

	num1 := int(a - '0')
	num2 := int(b - '0')
	num3 := int(c - '0')
	num4 := int(d - '0')

	num1 = ((num1 + 10) - 7) % 10
	num2 = ((num2 + 10) - 7) % 10
	num3 = ((num3 + 10) - 7) % 10
	num4 = ((num4 + 10) - 7) % 10

	decryptedValue := fmt.Sprintf("%d%d%d%d", num3, num4, num1, num2)

	fmt.Printf("\nThe decrypted value of %s is %s\n", number, decryptedValue)
}
