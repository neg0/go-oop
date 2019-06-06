package package_oriented_design

import (
	"github.com/neg0/go-oop/src/encapsulation/package_oriented_design"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCreditCard(t *testing.T) {
	sut := package_oriented_design.CreateCreditCard(
		"John Doe",
		"3233-3234-3242-3232",
		10,
		2022,
		752,
		1209.76,
	)

	assert.Equal(t, "John Doe", sut.OwnerName())
	assert.Equal(t, "3233-3234-3242-3232", sut.CardNumber())
	assert.Equal(t, 10, sut.ExpirationMonth())
	assert.Equal(t, 2022, sut.ExpirationYear())
	assert.Equal(t, 752, sut.SecurityCode())
	assert.Equal(t, float32(1209.76), sut.AvailableCredit())
}
