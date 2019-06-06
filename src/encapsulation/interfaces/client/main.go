package main

import (
	"github.com/neg0/go-oop/src/encapsulation/interfaces/payment"
)

func main() {
	var payOption payment.ProcessOption

	cashOpt := payOption.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	println(cashOpt.OwnerName())
}
