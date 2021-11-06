package main

import (
	"github.com/poncos/go-things/kafka-samples/internal/kafka"
)

func main() {

	kafkaClient := kafka.NewKafkaClient("localhost:29092,localhost:39092,localhost:49092", "myTopic.v2")
	defer kafkaClient.Close()

	var msg = map[string]interface{} {
		"originatingNumber": "123456789123456",
		"terminatingNumber": "123456789123456",
		"startingTime": "2021-11-02",
		"duration": int32(60),
		"callType": "Data",
	}

	kafkaClient.Publish(msg)
	kafkaClient.Flush(15 * 1000)
}