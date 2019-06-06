package pay

import "github.com/neg0/go-oop/src/polymorphism/payment"

type Pay struct{}

var option payment.Options

func (pp *Pay) WithDebitCard(amount float32) string {
	option = new(payment.DebitCard)
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCreditCard(amount float32) string {
	option = new(payment.CreditCard)
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCash(amount float32) string {
	option = new(payment.Cash)
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCryptoCurrency(amount float32) string {
	option = new(payment.CryptoCurrency)
	return option.ProcessPayment(amount)
}
