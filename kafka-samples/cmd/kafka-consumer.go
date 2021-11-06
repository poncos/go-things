package main

import (
	"fmt"
	"github.com/poncos/go-things/kafka-samples/internal/kafka"
)


func main() {

	kafkaClient := kafka.NewKafkaClient("localhost:29092,localhost:39092,localhost:49092", "myTopic.v2")
	defer kafkaClient.Close()

	msg, err  := kafkaClient.Consume()
	
	if (err == nil) {
		fmt.Printf("Received msg: [%v]\n", msg)
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}
}
