package cypher

const firstASCIICode = byte(97)

const lastASCIICode = byte(122)

const numberOfCharacters = lastASCIICode - firstASCIICode + 1

// Cypher encrypts the given text using the caesar algorithm
func Cypher(text string, key byte) string {
	result := make([]byte, len(text))

	for k, v := range text {
		if v != ' ' {
			result[k] = nextASCIICode(byte(v), key)
		} else {
			result[k] = byte(v)
		}
	}

	return string(result[:])
}

func nextASCIICode(code byte, key byte) byte {
	offset := (code - firstASCIICode + key) % numberOfCharacters
	nextCode := firstASCIICode + offset

	return nextCode
}
