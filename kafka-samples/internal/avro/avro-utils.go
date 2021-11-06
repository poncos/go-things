
import (
	"fmt"
)

const smechaDir = "conf"

func encode(obj interface {}) []byte {

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

	fmt.Printf("Value [%t]\n", record)

	return buffer.Bytes()
}