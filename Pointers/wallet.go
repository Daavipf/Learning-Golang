package pointers

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

// Withdraws some amount to the wallet's balance
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
