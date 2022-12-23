package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWithDrawFromCreditCardReturnsTrue(t *testing.T) {
	creditCardAccount := CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", cardNumber: "123456789"}

	result := creditCardAccount.CanWithDraw(300)

	assert.Equal(t, true, result)
}

func TestCanWithDrawFromCreditCardReturnsFalse(t *testing.T) {
	creditCardAccount := CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District"}

	result := creditCardAccount.CanWithDraw(1300)

	assert.Equal(t, false, result)
}

func TestWithDrawFromCreditCard(t *testing.T) {
	creditCardAccount := CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District"}

	creditCardAccount.WithDraw(500)

	assert.Equal(t, float64(500), creditCardAccount.GetBalance())
}

func TestGetIdentifierFromCreditCard(t *testing.T) {
	creditCardAccount := CreditCardAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", accountType: "accountType"}
	expectedIdentifier := "accountType"

	assert.Equal(t, expectedIdentifier, creditCardAccount.GetIdentifier())
}
