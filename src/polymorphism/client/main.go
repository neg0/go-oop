package main

import "github.com/neg0/go-oop/polymorphism/pay"

var paymentMethod *pay.Pay

func init() {
	paymentMethod = &pay.Pay{}
}

func main() {
	paymentMethod.WithCash(10)
	paymentMethod.WithCreditCard(120.95)
	paymentMethod.WithDebitCard(201.23)
	paymentMethod.WithCryptoCurrency(201.23)
}
