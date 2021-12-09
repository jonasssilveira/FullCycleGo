package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, ExpirationCVV int) (*CreditCard, error) {

	credit := &CreditCard{number, name, expirationMonth, expirationYear, ExpirationCVV}

	err := credit.isValid()

	if err != nil {
		return nil, err
	}

	return credit, nil

}

func (c *CreditCard) isValid() error {
	var erros []error
	erros = append(erros, c.checkYear())
	erros = append(erros, c.checkNumber())
	erros = append(erros, c.checkMonth())

	for _, erro := range erros {
		if erro != nil {
			return erro
		}
	}
	return nil
}

func (c *CreditCard) checkMonth() error {
	if c.expirationMonth > 0 && c.expirationMonth < 13 {
		return nil
	}
	return errors.New("Cartao com mes invalido")
}

func (c *CreditCard) checkYear() error {
	if c.expirationYear < time.Now().Year() {
		return errors.New("Cartao vencido")
	}
	return nil
}

func (c *CreditCard) checkNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.number) {
		return errors.New("Numero de cartao invalido")
	}
	return nil
}
