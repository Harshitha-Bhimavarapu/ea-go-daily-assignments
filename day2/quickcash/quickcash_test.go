package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavings(t *testing.T) {
	savingsAccount := &SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", accountType: "savings"}
	creditCardAccount := &CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", cardNumber: "123456789", accountType: "creditCard"}
	quickCash := QuickCash{savingsAccount, creditCardAccount}

	balanceAmount, accType := quickCash.getCash(500)

	assert.Equal(t, float64(500), balanceAmount)
	assert.Equal(t, "savings", accType)
}

func TestGetCashFromCreditCard(t *testing.T) {
	savingsAccount := &SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", accountType: "savings"}
	creditCardAccount := &CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 10000, address: "Street1, District", cardNumber: "123456789", accountType: "creditCard"}
	quickCash := QuickCash{savingsAccount, creditCardAccount}

	balanceAmount, accType := quickCash.getCash(5000)

	assert.Equal(t, float64(5000), balanceAmount)
	assert.Equal(t, "creditCard", accType)
}
