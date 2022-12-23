package account

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	err := acc.Withdraw(200)

	if err != nil {
		t.Fatalf("Insufficient Balance")
	}
	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestWithDrawFailsOnInsufficientBalance(t *testing.T) {
	balance := float64(5000)
	amount := float64(6000)
	acc := Account{balance: balance}
	errMessage := fmt.Sprintf("your current balance is Rs.%f. you are lacking Rs.%f to withdraw Rs.%f", balance, (amount - balance), amount)
	expectedError := InsufficientBalanceError{errMessage}

	err := acc.Withdraw(amount)

	if err != expectedError {
		t.Fatalf("Insufficient Balance")
	}
}
