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

func (kd *kafkaClient) consume(msg interface{}) {
/*
* c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092,localhost:49092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic.v2"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			avro_decode(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
*/
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