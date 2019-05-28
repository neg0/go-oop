package credit_account

import "github.com/neg0/go-oop/composition/account"

type CreditAccount struct{
	account.Account
}

func (ca *CreditAccount) RequestCreditIncrease() bool {
	return true
}