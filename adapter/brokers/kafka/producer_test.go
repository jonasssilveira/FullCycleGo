package kafka

import (
	"FullCycle/adapter/brokers/kafka/transaction"
	"FullCycle/domain/entity"
	"FullCycle/usecase/proccess_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducerPublish(t *testing.T) {
	expectOutput := proccess_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECT,
		ErrorMessage: "you dont have limit enough",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
