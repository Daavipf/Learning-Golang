package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("test deposit to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw to wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("test withdraw more than wallet's amount", func(t *testing.T) {
		startingBalance := Bitcoin(30)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(50))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, expected Bitcoin) {
	actual := w.Balance()

	if actual != expected {
		t.Errorf("actual %s expected %s", actual, expected)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
