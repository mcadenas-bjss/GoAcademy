package pointersanderrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient funds")

type Bitcoin float64

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount // automatically dereferenced
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance // dereferencing the pointer to get the value
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
