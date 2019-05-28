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

func TestPay_WithCash(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Cash")
}

func TestPay_WithCreditCard(t *testing.T) {
	assert.Contains(t, sut.WithCreditCard(10), "Creditcard")
}

func TestPay_WithDebitCard(t *testing.T) {
	assert.Contains(t, sut.WithDebitCard(10), "Debitcard")
}

func TestPay_WithCryptoCurrency(t *testing.T) {
	assert.Contains(t, sut.WithCryptoCurrency(10), "Crypto")
}
