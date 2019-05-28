package pay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sut *Pay

func init() {
	sut = new(Pay)
}

func TestPay_WithCash(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Cash")
}

func TestPay_WithCreditCard(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Creditcard")
}

func TestPay_WithDebitCard(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Debitcard")
}

func TestPay_WithCryptoCurrency(t *testing.T) {
	assert.Contains(t, sut.WithCash(10), "Crypto")
}
