package json

import (
	"encoding/json"
	"log"
)

type topic1 struct {
	// valid JSON encoding
}

// A. 'true' is not JSON encoding
// B. 'null' is not JSON encoding
// C. panic
// D. no output
func (*topic1) t1() {
	if !json.Valid([]byte("true")) {
		log.Fatal("'true' is not JSON encoding")
	}
	if !json.Valid([]byte("null")) {
		log.Println("'null' is not JSON encoding")
	}
}
