package brokers

type ProducerInterface interface {
	Publish(msg interface{}, key []byte, topic string) error
}
