package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000", "Jonas da Silva Silveira", 05, 20212, 123)
	assert.Equal(t, "Numero de cartao invalido", err.Error())
	_, err = NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 05, 2022, 123)
	assert.Nil(t, err)

}

func TestCreditCardMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 15, 2022, 123)
	assert.Equal(t, "Cartao com mes invalido", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 05, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardYear(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 15, 2015, 123)
	assert.Equal(t, "Cartao vencido", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jonas da Silva Silveira", 05, 2022, 123)
	assert.Nil(t, err)
}
