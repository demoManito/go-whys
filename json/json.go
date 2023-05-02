package json

import (
	"encoding/json"
	"log"
)

// Topic1 .
// A. 'true' is not JSON encoding
// B. 'null' is not JSON encoding
// C. panic
// D. no output
func Topic1() {
	if !json.Valid([]byte("true")) {
		log.Fatal("'true' is not JSON encoding")
	}
	if !json.Valid([]byte("null")) {
		println("'null' is not JSON encoding")
	}
}
