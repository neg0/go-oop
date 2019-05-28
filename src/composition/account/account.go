package account

type Account struct{}

func (a *Account) AvailableFunds() float32 {
	return 23.20
}

func (a *Account) ProcessPayment(amount float32) bool {
	return true
}
