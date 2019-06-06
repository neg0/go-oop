package credit_account

import "fmt"

type CreditAccount struct{
	balance float32
}

func CreateCreditAccount(chargeCh chan float32, initialBalance float32) *CreditAccount {
	creditAccount := &CreditAccount{balance: initialBalance}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

func (c *CreditAccount) processPayment(amount float32) {
	c.balance = c.balance - amount
	fmt.Println(fmt.Sprintf("New Balance after process of payment is: %.2f",c.balance))
}
