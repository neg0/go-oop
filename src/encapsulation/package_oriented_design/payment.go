package package_oriented_design

type CreditCard struct {
	ownerName string
	cardNumber string
	expirationMonth int
	expirationYear int
	securityCode int
	availableCredit float32
}

func CreateCreditCard(
	ownerName string,
	cardNumber string,
	expirationMonth int,
	expirationYear int,
	securityCode int,
	availableCredit float32,
	) *CreditCard {
	return &CreditCard{
		ownerName: ownerName,
		cardNumber: cardNumber,
		expirationMonth: expirationMonth,
		expirationYear: expirationYear,
		securityCode: securityCode,
		availableCredit: availableCredit,
	}
}

func (cc *CreditCard) OwnerName() string {
	return cc.ownerName
}

func (cc *CreditCard) CardNumber() string {
	return cc.cardNumber
}

func (cc *CreditCard) ExpirationMonth() int {
	return cc.expirationMonth
}

func (cc *CreditCard) ExpirationYear() int {
	return  cc.expirationYear
}

func (cc *CreditCard) SecurityCode() int {
	return cc.securityCode
}

func (cc *CreditCard) AvailableCredit() float32 {
	return cc.availableCredit
}
