package main

import (
	"errors"
	"testing"
)

/*
// Improve 1
func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)
	want := 10.0
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	t.Run("Sufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		wallet.Withdraw(5.0)
		want := 5.0
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("Insufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		err := wallet.Withdraw(15.0)
		if err == nil {
			t.Errorf("Insufficient funds on wallet.")
		}
	})
}
*/

/*
// Improve 2. El error que obtengo es el error esperado
func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)
	want := 10.0
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	t.Run("Sufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		wallet.Withdraw(5.0)
		want := 5.0
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("Insufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		err := wallet.Withdraw(15.0)
		if err == nil {
			t.Errorf("Error should be ocurred.")
		}
		if err != ErrInsufficientfunds {
			t.Errorf("Error should be %q but got %q", ErrInsufficientfunds, err)
		}
	})
}
*/

// Improve 3
func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)
	want := 10.0
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	t.Run("Sufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		wallet.Withdraw(5.0)
		want := 5.0
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("Insufficient cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		err := wallet.Withdraw(15.0)
		if err == nil {
			t.Errorf("Error should be ocurred.")
		}
		var got ErrInsufficientfunds
		isErrInsufficientfunds := errors.As(err, &got)
		if !isErrInsufficientfunds {
			t.Fatalf("was not a ErrInsufficientfunds, got %q", err)
		}
	})
}
