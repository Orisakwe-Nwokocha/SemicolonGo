package chapter4

import (
	"fmt"
	"semicolonGo/utils"
	"strings"
)

func CreditLimit() {
	input := utils.Input{}

	for {
		fmt.Println("Do you want to enter a customer account: y/n?: ")
		choice := input.NextLine()
		choice = strings.ToLower(choice)

		if choice != "y" && choice != "n" {
			println("Invalid input. Try again.")
			continue
		}

		if choice == "n" {
			break
		}

		fmt.Println("Enter customer account number: ")
		accountNumber := input.NextLine()

		fmt.Println("Enter customer initial balance: ")
		balance, _ := input.NextInt()

		fmt.Println("Enter customer total charges: ")
		totalCharges, _ := input.NextInt()

		fmt.Println("Enter customer total credits: ")
		totalCredits, _ := input.NextInt()

		fmt.Println("Enter customer allowed credit limit: ")
		allowedCreditLimit, _ := input.NextInt()

		newBalance := (balance + totalCharges) - totalCredits
		fmt.Printf("\nCustomer %s new balance: %d\n\n", accountNumber, newBalance)

		if newBalance > allowedCreditLimit {
			fmt.Printf("Credit limit exceeded for customer %s\n\n", accountNumber)
		}
	}
}
