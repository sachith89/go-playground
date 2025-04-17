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

	ebt, profit, ratio := calculateProfit(revenue, expense, taxRate)

	formattedTaxRate := fmt.Sprintf("Earnings Before Tax to Profit After Tax Ratio: %.2f\n", ratio)

	outputText(ebt, profit, formattedTaxRate)
}

func outputText(ebt float64, profit float64, formattedTaxRate string) {
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
	fmt.Printf(formattedTaxRate)
}

func calculateProfit(revenue float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {

	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
