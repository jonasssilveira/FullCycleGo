package proccess_transaction

import (
	"FullCycle/adapter/brokers"
	"FullCycle/domain/entity"
	"FullCycle/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer brokers.ProducerInterface
	Topic string
}

func NewProcessTransaction(repository repository.TransactionRepository, producer brokers.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
		Topic: topic,
		Producer: producer,
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

	output :=  TransactionDtoOutput{transaction.ID,
		entity.APPROVED, ""}
	err = p.publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transactional, invalidCC error) (TransactionDtoOutput, error) {

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, entity.REJECT, invalidCC.Error(),
		transaction.Amount)

	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output :=  TransactionDtoOutput{transaction.ID,
		entity.REJECT, invalidCC.Error()}
 	err = p.publish(output, []byte(transaction.ID))
	return output, nil

}

func (p *ProcessTransaction) publish(output TransactionDtoOutput, key []byte) error{
	err := p.Producer.Publish(output, key, p.Topic)

	if err != nil{
		return err
	}
	return nil
}