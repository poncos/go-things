package main

import (
	"fmt"
	"bytes"
	"os"
	"log"
	"gopkg.in/avro.v0"
	"path/filepath"
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
	return buffer.Bytes()
}

func main() {

	avro_msg := avro_encode("123456789123456","123456789123456","2021-11-02")
	fmt.Printf("Avro msg [%s]\n", string(avro_msg))
}