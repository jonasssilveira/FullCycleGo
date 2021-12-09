package proccess_transaction

import (
	"FullCycle/domain/entity"
	"FullCycle/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
	}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {

	transaction := entity.NewTransaction(input.ID, input.AccountID, "", "", input.Amount)

	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber,
		input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear,
		input.CreditCardCVV)

	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}
	transaction.SetCredicard(cc)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entity.Transactional) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, entity.APPROVED, "",
		transaction.Amount)
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	return TransactionDtoOutput{transaction.ID,
		entity.APPROVED, ""}, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transactional, invalidCC error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, entity.REJECT, invalidCC.Error(),
		transaction.Amount)
	if err != nil {
		return TransactionDtoOutput{}, err
	}
	return TransactionDtoOutput{transaction.ID,
		entity.REJECT, invalidCC.Error()}, nil
}
