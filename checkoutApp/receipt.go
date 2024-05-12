package checkoutapp

import (
	"fmt"
	"time"
)

func displayReceipt() {
	customerName, items, purchasedQuantities, itemPrices, cashierName,
		subTotal, discount, VAT, billTotal := displayInvoice()
	display("")
	amountPaid := getAmountPaid(billTotal)
	balance := amountPaid - billTotal

	headerFormat := `
SEMICOLON STORES
MAIN BRANCH
LOCATION: 312 HERBERT MACAULAY WAY, SABO YABA, LAGOS.
TEL: 03293828343
Date: %s
Cashier: %s
Customer Name: %s`
	fmt.Printf(headerFormat, time.Now().Format("02-Jan-06 3:04:05 pm"), cashierName, customerName)

	fmt.Print(`
======================================================
       ITEM          QTY       PRICE       TOTAL(NGN)
------------------------------------------------------
`)

	format1 := `   %10s       %3d    %10.2f     %10.2f
`
	for index := 0; index < len(items); index++ {
		fmt.Printf(format1, items[index], purchasedQuantities[index], itemPrices[index],
			itemPrices[index]*float64(purchasedQuantities[index]))
	}

	format2 := `-------------------------------------------------------
				Sub Total:        %10.2f
				 Discount:        %10.2f
			  VAT @ 7.50%%:        %10.2f
=======================================================
			   Bill Total:        %10.2f
			  Amount Paid:        %10.2f
			      Balance:        %10.2f
=======================================================
		THANK YOU FOR YOUR PATRONAGE
=======================================================
`
	fmt.Printf(format2, subTotal, discount, VAT, billTotal, amountPaid, balance)
}
