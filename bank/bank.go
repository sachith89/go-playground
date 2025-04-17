package main

import "fmt"

const divider = "-----------------------------------"

func main() {

	var accountBalance = 1000.0

	fmt.Printf("\n%v\nWelcome to Go banking app\n%v\n", divider, divider)

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice uint

	fmt.Print("\nEnter your choice: ")
	_, _ = fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
		return
	} else if choice == 2 {
		fmt.Print("Enter the amount to deposit: ")
		var depositAmount float64
		_, _ = fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your balance is:", accountBalance)
		return
	} else if choice == 3 {
		fmt.Print("Enter the amount to withdraw: ")
		var withdrawAmount float64
		_, _ = fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Your balance is:", accountBalance)
		}
		return
	}

	fmt.Println("Invalid choice")
}
