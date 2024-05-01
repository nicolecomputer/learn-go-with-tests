package poe

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet *Wallet, balance Bitcoin) {
		t.Helper()
		if wallet.balance != balance {
			t.Errorf("got %s, expected %s", wallet.balance, balance)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))
		if err != nil {
			t.Errorf("Received unexpected error: %s", err)
		}

		assertBalance(t, &wallet, Bitcoin(5))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(1))
		err := wallet.Withdraw(Bitcoin(5))

		if err != ErrInsufficientFunds {
			t.Errorf("Expected an error when overdrawing")
		}
		assertBalance(t, &wallet, Bitcoin(1))
	})
}
