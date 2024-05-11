package checkoutapp

import (
	"fmt"
	"semicolonGo/utils"
	"time"
)

func Checkout() {

	displayInvoice()

}

func displayInvoice() {
	customerName := input("What is the customer's name?")
	items, purchasedQuantities, itemPrices := collectCheckoutDetails()
	cashierName := input("What is your name?")
	subTotal := getSubTotal(purchasedQuantities, itemPrices)
	discount := (getDiscount() / 100) * subTotal
	VAT := subTotal * (17.50 / 100)
	billTotal := (subTotal - discount) + VAT

	format := `
SEMICOLON STORES
MAIN BRANCH
LOCATION: 312 HERBERT MACAULAY WAY, SABO YABA, LAGOS.
TEL: 03293828343
Date: %s
Cashier: %s
Customer Name: %s
`
	fmt.Printf(format, time.Now().Format("02-Jan-06 at 3:04:05 PM"), cashierName, customerName)

	fmt.Println("======================================================")
	fmt.Println("       ITEM          QTY     PRICE        TOTAL(NGN)")
	fmt.Println("------------------------------------------------------")

	for index := 0; index < len(items); index++ {
		fmt.Printf("   %10s", items[index])
		fmt.Printf("         %3d", purchasedQuantities[index])
		fmt.Printf("  %10.2f", itemPrices[index])
		fmt.Printf("     %10.2f\n", itemPrices[index]*float64(purchasedQuantities[index]))
	}

	fmt.Printf("\n\n-------------------------------------------------------\n")
	fmt.Printf("\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n",
		"Sub Total:", subTotal, "Discount:", discount, "VAT @ 7.50%:", VAT)
	fmt.Printf("=======================================================\n")
	fmt.Printf("\t\t\t%13s\t  %10.2f\n", "Bill Total:", billTotal)
	fmt.Printf("=======================================================\n")
	fmt.Printf("   THIS IS NOT A RECEIPT, KINDLY PAY %.2f\n", billTotal)
	fmt.Printf("=======================================================\n")
}

func getSubTotal(purchasedQuantities []int, itemPrices []float64) float64 {
	sum := 0.0
	for index, itemPrice := range itemPrices {
		sum += itemPrice * float64(purchasedQuantities[index])
	}
	return sum
}

func input(prompt string) string {
	scanner := utils.Input{}
	display(prompt)

	return scanner.NextLine()
}

func display(output any) {
	fmt.Println(output)
}
