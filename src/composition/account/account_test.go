package account

import (
	"testing"

	"github.com/neg0/go-oop/composition/account"
	"github.com/stretchr/testify/assert"
)

var sut *account.Account

func init() {
	sut = new(account.Account)
}

func TestAccount_AvailableFunds(t *testing.T) {
	assert.Equal(t, float32(23.20), sut.AvailableFunds())
}

func TestAccount_ProcessPayment(t *testing.T) {
	assert.True(t, sut.ProcessPayment(12000))
}