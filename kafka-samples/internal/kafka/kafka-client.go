package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/poncos/go-things/kafka-samples/internal/avro"
)

type kafkaClient struct {
	kafkaProducer *kafka.Producer
	kafkaConsumer *kafka.Consumer
	bootstrapServers string
	topicName string
}

func NewKafkaClient(bootstrapServers string, topic string) *kafkaClient {
	
	obj := new(kafkaClient)
    obj.bootstrapServers = bootstrapServers
	obj.topicName = topic

    return obj
}

func (kc *kafkaClient) Close() {
	if (kc.kafkaProducer != nil) {
		kc.kafkaProducer.Close()
	}

	if (kc.kafkaConsumer != nil) {
		kc.kafkaConsumer.Close()
	}
}

func (kc *kafkaClient) Publish(msg map[string]interface{}) {

	if (kc.kafkaProducer == nil) {
		kc.createProducer()
	}

	avro_msg := avro.EncodeBinary(msg)
	kc.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kc.topicName, Partition: kafka.PartitionAny},
		Value: avro_msg,
		}, nil )
}

func  (kc *kafkaClient) Flush(timeoutMS int) {
	if (kc.kafkaProducer != nil) {
		kc.kafkaProducer.Flush(timeoutMS)
	}
}

func (kc *kafkaClient) Consume() (interface{}, error) {

	if (kc.kafkaConsumer == nil) {
		kc.createConsumer()
	}

	var err error
	msg, err := kc.kafkaConsumer.ReadMessage(-1)
	if err == nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		return avro.DecodeBinary(msg.Value), nil
	} else {
		// The client will automatically try to recover from all errors.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		return nil, err
	}
}

func (kc *kafkaClient) createProducer() {

	var err error
	kc.kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092,localhost:39092,localhost:49092"})
	if err != nil {
		panic(err)
	}

	// go-rutine: Delivery report handler for produced messages
	go func() {
		for e := range kc.kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %t\n", ev.TopicPartition)
				}
			}
		}
	}()
}

func (kc *kafkaClient) createConsumer() {

	var err error
	kc.kafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kc.bootstrapServers,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kc.kafkaConsumer.SubscribeTopics([]string{kc.topicName}, nil)
}