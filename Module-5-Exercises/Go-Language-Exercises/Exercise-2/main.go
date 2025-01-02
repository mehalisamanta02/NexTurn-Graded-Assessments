package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      int
	Name    string
	Balance float64
	History []string
}

var accounts []Account

// Deposit money into account
func deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accountID {
			accounts[i].Balance += amount
			accounts[i].History = append(accounts[i].History, fmt.Sprintf("Deposited: Rs.%.2f", amount))
			return nil
		}
	}
	return errors.New("Account not found")
}

// Withdraw money from account
func withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accountID {
			if acc.Balance < amount {
				return errors.New("insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].History = append(accounts[i].History, fmt.Sprintf("Withdrew: Rs.%.2f", amount))
			return nil
		}
	}
	return errors.New("Account not found")
}

// View transaction history
func viewHistory(accountID int) ([]string, error) {
	for _, acc := range accounts {
		if acc.ID == accountID {
			return acc.History, nil
		}
	}
	return nil, errors.New("Account not found")
}

// Menu system
func Menu() {
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&amount)
			if err := deposit(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case 2:
			var id int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scanln(&amount)
			if err := withdraw(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case 3:
			var id int
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&id)
			for _, acc := range accounts {
				if acc.ID == id {
					fmt.Printf("Account Balance: Rs.%.2f\n", acc.Balance)
				}
			}

		case 4:
			var id int
			fmt.Print("Enter account ID: ")
			fmt.Scanln(&id)
			if history, err := viewHistory(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, entry := range history {
					fmt.Println(entry)
				}
			}

		case 5:
			fmt.Println("Exited.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {

	accounts = append(accounts, Account{ID: 1234, Name: "Rajiv Shaw", Balance: 5000.0})
	accounts = append(accounts, Account{ID: 2234, Name: "Sheetal Chhabra", Balance: 25000.0})
	accounts = append(accounts, Account{ID: 3234, Name: "Sameer Kapahi", Balance: 50000.0})
	Menu()
}
