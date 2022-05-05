package main

import "fmt"

type Account struct {
	name   string
	number int
	bank   string
}

func main() {
	var acc1 Account
	var accNumber int

	acc1.name = "Harman Kibue"
	acc1.number = 123456789
	acc1.bank = "NCBA"

	accNumber = printAccount(acc1)
	fmt.Println(accNumber)

}

func printAccount(acc Account) int {
	fmt.Println("The ACC name is: ", acc.name)
	fmt.Println("The ACC number is:", acc.number)
	fmt.Println("The ACC bank is:", acc.bank)
	fmt.Println("-------------------")

	return acc.number
}
