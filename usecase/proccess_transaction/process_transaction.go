package proccess_transaction

import (
	"FullCycle/adapter/brokers"
	"FullCycle/domain/entity"
	"FullCycle/domain/repository"
	"log"
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
	log.Println("Vai executar o ",input)

	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber,
		input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear,
		input.CreditCardCVV)
	log.Println("Invalido? ",cc,invalidCC)

	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}
	transaction.SetCredicard(cc)
	invalidTransaction := transaction.IsValid()
	log.Println("Executou: ",transaction)
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
	log.Println("Aprovou: ",output)

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transactional, invalidCC error) (TransactionDtoOutput, error) {
	log.Println("Vai rejeitar: ",transaction)

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, entity.REJECT, invalidCC.Error(),
		transaction.Amount)
	log.Println("salvou? ",err)

	if err != nil {
		return TransactionDtoOutput{}, err
	}
	output :=  TransactionDtoOutput{transaction.ID,
		entity.REJECT, invalidCC.Error()}
	log.Println("Rejeitou: ",output)
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