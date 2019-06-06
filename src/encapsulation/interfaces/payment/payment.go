package payment

import "fmt"

type ProcessOption interface {
	ProcessEncapsulatedPayment(amount float32) bool
}

type CreditAccountDefinition interface {
	OwnerName() string
	CardNumber() string
	ExpirationMonth() int
	ExpirationYear() int
	SecurityCode() int
	AvailableCredit() float32
}

type CreditAccount struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func CreateCreditAccount(
	ownerName string,
	cardNumber string,
	expirationMonth int,
	expirationYear int,
	securityCode int,
	availableCredit float32) *CreditAccount {
	return &CreditAccount{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
		availableCredit: availableCredit,
	}
}

func (ca *CreditAccount) ProcessEncapsulatedPayment(amount float32) bool {
	fmt.Println("Processing a Credit card payment...")
	return true
}

func (ca CreditAccount) OwnerName() string {
	return ca.ownerName
}

func (ca *CreditAccount) SetOwnerName(ownerName string) {
	ca.ownerName = ownerName
}

func (ca CreditAccount) CardNumber() string {
	return ca.cardNumber
}

func (ca *CreditAccount) SetCardNumber(cardNumber string) {
	ca.cardNumber = cardNumber
}

func (ca *CreditAccount) ExpirationMonth() int {
	return ca.expirationMonth
}

func (ca *CreditAccount) SetExpirationMonth(expirationMonth int)  {
	ca.expirationMonth = expirationMonth
}

func (ca CreditAccount) ExpirationYear() int {
	return ca.expirationYear
}

func (ca *CreditAccount) SetExpirationYear(expirationYear int) {
	ca.expirationYear = expirationYear
}

func (ca CreditAccount) SecurityCode() int {
	return ca.securityCode
}

func (ca *CreditAccount) SetSecurityCode(securityCode int) {
	ca.securityCode = securityCode
}

func (ca CreditAccount) AvailableCredit() float32 {
	return ca.availableCredit
}

func (ca *CreditAccount) SetAvailableCredit(availableCredit float32) {
	ca.availableCredit = availableCredit
}
