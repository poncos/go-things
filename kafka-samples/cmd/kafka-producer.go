package main

import (
	"fmt"
	"bytes"
	"os"
	"log"
	"gopkg.in/avro.v0"
	"path/filepath"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const schemaFilePath string = "conf/cdr-sample.avsc"

func avro_encode(originatingNumber string, terminatingNumber string, startingTime string) []byte {

	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	exPath := filepath.Dir(ex)

	schemaFilePath := fmt.Sprintf("%s/%s", exPath, schemaFilePath)
	schema, err := avro.ParseSchemaFile(schemaFilePath)

	record := avro.NewGenericRecord(schema)
	record.Set("originatingNumber", originatingNumber)
	record.Set("terminatingNumber", terminatingNumber)
	record.Set("startingTime", startingTime)

	writer := avro.NewGenericDatumWriter()
	// SetSchema must be called before calling Write
	writer.SetSchema(schema)

	// Create a new Buffer and Encoder to write to this Buffer
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	// Write the record
	err = writer.Write(record, encoder)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Avro msg [%s]\n", string(buffer.Bytes()))
	fmt.Printf("Value [%t]\n", record)

	return buffer.Bytes()
}

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092,localhost:39092,localhost:49092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
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

	// Produce messages to topic (asynchronously)
	topic := "myTopic.v2"
	for i := 0; i<10; i++ {

		avro_msg := avro_encode("asfdasfdasfasfdadsf","123456789123456","2021-11-02")
		fmt.Printf("Publishing msg [%s]\n", string(avro_msg))
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          avro_msg,
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}