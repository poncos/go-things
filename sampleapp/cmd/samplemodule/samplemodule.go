package main

import (
	"fmt"

	"github.com/poncos/go-things/samplemodule/pkg/cypher"
)

func main() {
	text := "hello how are you"

	fmt.Printf("%s\n", text)

	cyphered := cypher.Cypher(text, 3)
	fmt.Printf("Cyphered text: %s\n", cyphered)
}
