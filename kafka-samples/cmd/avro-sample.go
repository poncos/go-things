package main

import (
	"fmt"
	
	"github.com/poncos/go-things/kafka-samples/internal/avro"
)

func main() {

	var raw_msg = map[string]interface{} {
		"originatingNumber": "123456789123456",
		"terminatingNumber": "123456789123456",
		"startingTime": "2021-11-02",
		"duration": int32(60),
		"callType": "Data",
	}

	avro_msg := avro.EncodeBinary(raw_msg)
	fmt.Printf("Avro msg [%s]\n", string(avro_msg))

	decoded_msg := avro.DecodeBinary(avro_msg)
	fmt.Printf("Decoded msg [%s]\n", decoded_msg)
}