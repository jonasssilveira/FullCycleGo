package transaction

import (
	"encoding/json"

	"FullCycle/usecase/proccess_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMEssage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (kafka *KafkaPresenter) Bind(input interface{}) error {
	kafka.ID = input.(proccess_transaction.TransactionDtoOutput).ID
	kafka.Status = input.(proccess_transaction.TransactionDtoOutput).Status
	kafka.ErrorMEssage = input.(proccess_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (kafka *KafkaPresenter) Show() ([]byte, error) {

	j, err := json.Marshal(kafka)

	if err != nil {
		return nil, err
	}
	return j, nil
}
