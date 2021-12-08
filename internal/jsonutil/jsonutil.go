package jsonutil

import (
	"encoding/json"
	"opa-echo-test/internal/chk"
)

// Marshal .
func Marshal(obj interface{}) string {
	b, err := json.Marshal(obj)
	chk.SE(err)
	return string(b)
}

// Unmarshal .
func Unmarshal(str string, out interface{}) {
	chk.SE(json.Unmarshal([]byte(str), out))
}
