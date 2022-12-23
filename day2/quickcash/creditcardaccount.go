package quickcash

type CreditCardAccount struct {
	firstName   string
	lastName    string
	phoneNumber string
	balance     float64
	address     string
	accountType string
	cardNumber  string
}

func (ca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return ca.balance >= amount
}

func (ca *CreditCardAccount) WithDraw(amount float64) {
	ca.balance -= amount
}

func (ca *CreditCardAccount) GetBalance() float64 {
	return ca.balance
}

func (ca *CreditCardAccount) GetIdentifier() string {
	return ca.accountType
}
