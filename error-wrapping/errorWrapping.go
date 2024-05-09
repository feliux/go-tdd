package main

import (
	// "errors"
	"fmt"
)

/*
// Improve 1
type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return errors.New("insufficient funds on wallet")
	}
	w.balance -= amount
	return nil
}
*/

/*
// Improve 2
type Wallet struct {
	balance float64
}

var ErrInsufficientfunds = errors.New("insufficient funds on wallet")

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return ErrInsufficientfunds
	}
	w.balance -= amount
	return nil
}
*/

// Improve 3
type Wallet struct {
	balance float64
}

type ErrInsufficientfunds struct {
	balance float64
	want    float64
	//err     error // encapsular errores y no exponer errores críticos del sistema
}

func (e ErrInsufficientfunds) Error() string {
	return fmt.Sprintf("can not withdraw, insufficient funds on wallet: want %.2f, balance %.2f", e.want, e.balance)
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return ErrInsufficientfunds{
			balance: w.balance,
			want:    amount,
		}
	}
	w.balance -= amount
	return nil
}
