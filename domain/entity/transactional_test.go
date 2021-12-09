package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAmountOfTransaction(t *testing.T) {
	var transaction = NewTransaction("123",
		"321", "", "", 1230.1)
	assert.Equal(t, "you dont have limit enough", transaction.IsValid().Error())
	transaction.Amount = 100.
	assert.Nil(t, transaction.IsValid())
	transaction.Amount = 0.
	assert.Equal(t, "you cannot transfer minus than 0", transaction.IsValid().Error())
}

func TestIfSetCreditCardIsWorking(t *testing.T) {
	var transaction = NewTransaction("123",
		"321", "", "", 1230.1)
	assert.Nil(t, transaction.CreditCard)
	creditCard, _ := NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 05, 2022, 123)
	transaction.SetCredicard(creditCard)
	assert.NotNil(t, transaction.CreditCard)

}
