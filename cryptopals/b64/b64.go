package b64

import (
	"encoding/base64"
)

func Encode(plaintext []byte) string {
	return base64.StdEncoding.EncodeToString(plaintext)
}

func Decode(encoded string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return bytes
}
