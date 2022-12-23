package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithDrawCash(t *testing.T) {
	savingsAccount := &SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", accountType: "Debit Card"}
	creditCardAccount := &CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", cardNumber: "123456789", accountType: "Credit Card"}
	quickCash := &QuickCash{savingsAccount, creditCardAccount}
	paytmWallet := PaytmWallet{quickcash: *quickCash}

	balanceAmnt, accountType := paytmWallet.quickcash.getCash(500)

	assert.Equal(t, float64(500), balanceAmnt)
	assert.Equal(t, "Debit Card", accountType)
}
