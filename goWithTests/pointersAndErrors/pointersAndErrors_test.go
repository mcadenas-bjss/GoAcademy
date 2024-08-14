package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10.0)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(10.0))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(30.0))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func ExampleWallet_Deposit() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	fmt.Println(wallet.Balance())
	// Output: 10.000000 BTC
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wallet.Balance())
	// Output: 10.000000 BTC
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
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

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
