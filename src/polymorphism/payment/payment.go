package payment

import "fmt"

type CreditCard struct{}
type DebitCard struct{}
type Cash struct{}
type CryptoCurrency struct{}

type Options interface {
	ProcessPayment(amount float32) string
}

func (d *DebitCard) ProcessPayment(amount float32) string {
	return fmt.Sprintf("Paying with Debitcard for amount of: %.2f\n", amount)
}

func (c *CreditCard) ProcessPayment(amount float32) string {
	return fmt.Sprintf("Paying with Creditcard for amount of: %.2f\n", amount)
}

func (c *Cash) ProcessPayment(amount float32) string {
	return fmt.Sprintf("Paying with Cash for amount of: %.2f\n", amount)
}

func (c *CryptoCurrency) ProcessPayment(amount float32) string {
	return fmt.Sprintf("Paying with Crypto currency for amount of: %.2f\n", amount)
}
