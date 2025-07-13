package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, expected Bitcoin) {
		actual := w.Balance()

		if actual != expected {
			t.Errorf("actual %s expected %s", actual, expected)
		}
	}
	t.Run("test deposit to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw to wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(20))
	})
}
