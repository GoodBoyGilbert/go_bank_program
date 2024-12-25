package main

import (
	"fmt"
	"os"
)

func main() {

	var accountBalance = 1000.0

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0.")
			} else {
				accountBalance += depositAmount
				fmt.Println("Deposit successful! New balance: ", accountBalance)
			}
		} else if choice == 3 {
			fmt.Print("Your Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			
			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than 0.")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance. You cannot withdraw more than your current balance.")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawal successful! New balance: ", accountBalance)
			}
		} else if choice == 4 {
			fmt.Println("Exiting the program. Thank you for banking with us!")
			os.Exit(0) 
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println() 
	}
}
