package pay

import "github.com/neg0/go-oop/polymorphism/payment"

type PaymentOptions interface {
	ProcessPayment(float32) bool
}

type Pay struct{}

var option PaymentOptions

func (pp *Pay) WithDebitCard(amount float32) {
	option = &payment.DebitCard{}
	option.ProcessPayment(amount)
}

func (pp *Pay) WithCreditCard(amount float32) {
	option = &payment.CreditCard{}
	option.ProcessPayment(amount)
}

func (pp *Pay) WithCash(amount float32) {
	option = &payment.Cash{}
	option.ProcessPayment(amount)
}

func (pp *Pay) WithCryptoCurrency(amount float32) {
	option = &payment.CryptoCurrency{}
	option.ProcessPayment(amount)
}
