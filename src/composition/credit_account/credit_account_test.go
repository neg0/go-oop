package credit_account

import (
	"testing"

	"github.com/neg0/go-oop/src/composition/credit_account"
	"github.com/stretchr/testify/assert"
)

var sut *credit_account.CreditAccount

func init() {
	sut = &credit_account.CreditAccount{}
}

func TestCreditAccount_RequestCreditIncrease(t *testing.T) {
	assert.True(t, sut.RequestCreditIncrease())
}
