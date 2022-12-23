package account

import (
	"fmt"
)

type Account struct {
	balance float64
}

type InsufficientBalanceError struct {
	message string
}

func (balError InsufficientBalanceError) Error() string {
	return balError.message
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	balance := acc.balance
	if balance < amount {
		errMessage := fmt.Sprintf("your current balance is Rs.%f. you are lacking Rs.%f to withdraw Rs.%f", balance, (amount - balance), amount)
		return InsufficientBalanceError{errMessage}
	}
	acc.balance -= amount
	return nil
}
