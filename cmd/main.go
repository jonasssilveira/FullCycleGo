package main

import (
	"FullCycle/adapter/brokers/kafka"
	"FullCycle/adapter/brokers/kafka/transaction"
	"FullCycle/adapter/factory"
	"FullCycle/usecase/proccess_transaction"
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	fmt.Println("Inicio...")
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapProducer, topics)
	go consumer.Consume(msgChan)

	usecase := proccess_transaction.NewProcessTransaction(&repository, producer, "transaction_result")
	log.Println(usecase)
	for msg := range msgChan {
		var input proccess_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		log.Println(input)
		usecase.Execute(input)
	}

}
