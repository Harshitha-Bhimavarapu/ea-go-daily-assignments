package quickcash

type SavingsAccount struct {
	firstName   string
	lastName    string
	phoneNumber string
	balance     float64
	address     string
	accountType string
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return sa.balance >= amount
}

func (sa *SavingsAccount) WithDraw(amount float64) {
	sa.balance -= amount
}

func (sa *SavingsAccount) GetBalance() float64 {
	return sa.balance
}

func (sa *SavingsAccount) GetIdentifier() string {
	return sa.accountType
}
