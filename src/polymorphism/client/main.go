package main

import (
	"fmt"
	"github.com/neg0/go-oop/polymorphism/pay"
)

var paymentMethod *pay.Pay

func init() {
	paymentMethod = &pay.Pay{}
}

func main() {
	cash := paymentMethod.WithCash(10)
	fmt.Println(cash)

	cc := paymentMethod.WithCreditCard(120.95)
	fmt.Println(cc)

	dc := paymentMethod.WithDebitCard(201.23)
	fmt.Println(dc)

	cryptoc := paymentMethod.WithCryptoCurrency(201.23)
	fmt.Println(cryptoc)
}
