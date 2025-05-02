package main

import (
	"fmt"
	"os"
)

const divider = "-----------------------------------"

func writeBalanceToFile(balance float64) {
	// Write the balance to the file
	balanceText := fmt.Sprintf("Balance: %.2f\n", balance)

	if err := os.WriteFile("balance.txt", []byte(balanceText), 0644); err != nil {
		fmt.Println("Error writing to file")
	}

}

func main() {

	var accountBalance = 1000.0

	for {
		fmt.Printf("\n%v\nWelcome to Go banking app\n%v\n", divider, divider)

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice uint

		fmt.Print("\nEnter your choice: ")
		_, _ = fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Println("Your balance is: ", accountBalance)
			}

		case 2:
			{
				fmt.Print("Enter the amount to deposit: ")

				var depositAmount float64
				_, _ = fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Should be greater than 0")
					continue
				}

				accountBalance += depositAmount

				fmt.Println("Your balance is:", accountBalance)
			}

		case 3:
			{
				fmt.Print("Enter the amount to withdraw: ")

				var withdrawAmount float64
				_, _ = fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount. Should be greater than 0")
					continue
				}

				if withdrawAmount > accountBalance {
					fmt.Println("Insufficient funds")
					continue
				}

				accountBalance -= withdrawAmount
				fmt.Println("Your balance is:", accountBalance)
			}
		default:
			{
				fmt.Println("Goodbye")
				fmt.Println("Thank you for using Go banking app")
				return
			}
		}
	}

}
