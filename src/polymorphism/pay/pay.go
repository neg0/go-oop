package pay

import "github.com/neg0/go-oop/polymorphism/payment"

type Pay struct{}

var option payment.Options

func (pp *Pay) WithDebitCard(amount float32) string {
	option = &payment.DebitCard{}
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCreditCard(amount float32) string {
	option = &payment.CreditCard{}
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCash(amount float32) string {
	option = &payment.Cash{}
	return option.ProcessPayment(amount)
}

func (pp *Pay) WithCryptoCurrency(amount float32) string {
	option = &payment.CryptoCurrency{}
	return option.ProcessPayment(amount)
}
