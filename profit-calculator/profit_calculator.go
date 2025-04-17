package main

import "fmt"

func main() {

	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Enter the revenue: ")
	_, _ = fmt.Scan(&revenue)

	fmt.Print("Enter the expense: ")
	_, _ = fmt.Scan(&expense)

	fmt.Print("Enter the tax rate (as a percentage): ")
	_, _ = fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	formattedTaxRate := fmt.Sprintf("Earnings Before Tax to Profit After Tax Ratio: %.2f\n", ratio)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
	fmt.Printf(formattedTaxRate)
}
