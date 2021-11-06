package main

import (
	"log"
	"fmt"
	"os"
	"gopkg.in/avro.v0"
	"path/filepath"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const schemaFilePath string = "conf/cdr-sample.avsc"

func avro_decode(message []byte) {

	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	exPath := filepath.Dir(ex)

	schemaFilePath := fmt.Sprintf("%s/%s", exPath, schemaFilePath)
	schema, err := avro.ParseSchemaFile(schemaFilePath)

	reader := avro.NewGenericDatumReader()
	// SetSchema must be called before calling Read
	reader.SetSchema(schema)

	// Create a new Decoder with a given buffer
	decoder := avro.NewBinaryDecoder(message)

	decodedRecord := avro.NewGenericRecord(schema)
	// Read data into given GenericRecord with a given Decoder. The first parameter to Read should be something to read into
	err = reader.Read(decodedRecord, decoder)
	if err != nil {
		panic(err)
	}

}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
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
}
