package pointers

import "errors"

type Wallet struct {
	balance Bitcoin
}

// Adds some amount to the wallet's balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Returns the total balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraws some amount to the wallet's balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
