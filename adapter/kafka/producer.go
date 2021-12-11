package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewKafkaProducer() *Producer {
	return &Producer{}
}

func (producer *Producer) Publish(msg interface{}, key []byte, topic string) {

}
