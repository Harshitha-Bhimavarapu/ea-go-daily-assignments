package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWithDrawReturnsTrue(t *testing.T) {
	savingsAccount := SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District"}

	result := savingsAccount.CanWithDraw(300)

	assert.Equal(t, true, result)
}

func TestCanWithDrawReturnsFalse(t *testing.T) {
	savingsAccount := SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District"}

	result := savingsAccount.CanWithDraw(1300)

	assert.Equal(t, false, result)
}

func TestWithDraw(t *testing.T) {
	savingsAccount := SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District"}

	savingsAccount.WithDraw(500)

	assert.Equal(t, float64(500), savingsAccount.GetBalance())
}

func TestGetIdentifier(t *testing.T) {
	savingsAccount := SavingsAccount{firstName: "FirstName", lastName: "LastName", phoneNumber: "1234567890", balance: 1000, address: "Street1, District", accountType: "accountType"}
	expectedIdentifier := "accountType"

	assert.Equal(t, expectedIdentifier, savingsAccount.GetIdentifier())
}
