package main

import "errors"

type WellsFargo struct {
	balance int
}

func newWellsFargo() *WellsFargo {
	return &WellsFargo{balance: 0}
}

func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WellsFargo) Withdraw(amount int) error {
	newBalance := w.balance - amount
	if newBalance <= 0 {
		return errors.New("Insuficient Balance in your bank account!")
	}

	w.balance = newBalance
	return nil
}

func (w *WellsFargo) GetBalance() int {
	return w.balance
}
