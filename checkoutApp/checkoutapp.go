package checkoutapp

import (
	"fmt"
	"semicolonGo/utils"
)

func Checkout() {
	displayInvoice()
}

func input(prompt string) string {
	scanner := utils.Input{}
	display(prompt)

	return scanner.NextLine()
}

func display(output any) {
	fmt.Println(output)
}
