package cypher_test

import (
	"fmt"
	"testing"

	"com.github/go-things/samplemodule/pkg/cypher"
)

func TestCaesarCypher(t *testing.T) {
	text := "hello how are you"

	cypheredText := cypher.Cypher(text, 3)
	fmt.Println(cypheredText)
	if cypheredText != "khoor krz duh brx" {
		t.Errorf("khoor krz duh brx; want %s", cypheredText)
	}
}
