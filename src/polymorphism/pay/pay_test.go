package pay

import (
	"testing"

	"github.com/neg0/go-oop/polymorphism/pay"
	"github.com/stretchr/testify/assert"
)

var sut *pay.Pay

func init() {
	sut = new(pay.Pay)
}

func TestPay_PayWithCash(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Cash")
}

func TestPay_PayWithCreditCard(t *testing.T) {
	assert.Contains(t, sut.WithCreditCard(10), "Creditcard")
}

func TestPay_PayWithDebitCard(t *testing.T) {
	assert.Contains(t, sut.WithDebitCard(10), "Debitcard")
}

func TestPay_PayWithCryptoCurrency(t *testing.T) {
	assert.Contains(t, sut.WithCryptoCurrency(10), "Crypto")
}
