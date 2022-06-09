package main

import "fmt"

type BankAccount interface {
	Deposit(amount int)
	Withdraw(amount int) error
	GetBalance() int
}

func main() {
	acc := newWellsFargo()
	acc.Deposit(1000)
	balance := acc.GetBalance()
	fmt.Println("Initial balance: ", balance)
	acc.Withdraw(500)
	finalBalance := acc.GetBalance()
	fmt.Println("The final balance: ", finalBalance)

}
