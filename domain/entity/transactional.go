package entity

import (
	"errors"
)

const (
	APPROVED = "aproved"
	REJECT   = "reject"
)

type Transactional struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   *CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction(id, account, status, errorMessage string, amount float64) *Transactional {
	return &Transactional{ID: id,
		Amount:       amount,
		AccountID:    account,
		Status:       status,
		ErrorMessage: errorMessage}
}

func (t *Transactional) IsValid() error {

	var erros []error

	erros = append(erros, t.underAllow())
	erros = append(erros, t.noLimit())

	for _, erro := range erros {
		if erro != nil {
			return erro
		}
	}
	return nil
}

func (t *Transactional) noLimit() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit enough")
	}
	return nil
}

func (t *Transactional) underAllow() error {
	if t.Amount < 1 {
		return errors.New("you cannot transfer minus than 0")
	}
	return nil
}

func (t *Transactional) SetCredicard(c *CreditCard) {
	t.CreditCard = c
}
