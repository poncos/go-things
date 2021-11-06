package avro

import (
	"fmt"
	"bytes"
	"os"
	"log"

	"gopkg.in/avro.v0"
	"path/filepath"
)

const schemaFilePath = "conf/cdr-sample.avsc"

func EncodeBinary(values map[string]interface{}) []byte {

	schema, _ := avro.ParseSchemaFile(schemaFileFullPath())
	record := avro.NewGenericRecord(schema)

	for k,v := range values {
		record.Set(k, v)
	}

	writer := avro.NewGenericDatumWriter()
	// SetSchema must be called before calling Write
	writer.SetSchema(schema)

	// Create a new Buffer and Encoder to write to this Buffer
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	// Write the record
	err := writer.Write(record, encoder)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func DecodeBinary(message []byte) interface{} {

	schema, _ := avro.ParseSchemaFile(schemaFileFullPath())
	reader := avro.NewGenericDatumReader()

	// SetSchema must be called before calling Read
	reader.SetSchema(schema)

	// Create a new Decoder with a given buffer
	decoder := avro.NewBinaryDecoder(message)

	decodedRecord := avro.NewGenericRecord(schema)
	// Read data into given GenericRecord with a given Decoder. The first parameter to Read should be something to read into
	err := reader.Read(decodedRecord, decoder)
	if err != nil {
		panic(err)
	}

	return decodedRecord
}

func schemaFileFullPath() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	schemaFileFullPath := fmt.Sprintf("%s/%s", exPath, schemaFilePath)
	return schemaFileFullPath
}