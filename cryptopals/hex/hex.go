package hex

import (
	xeh "encoding/hex"
)

func Encode(plaintext []byte) string {
	return xeh.EncodeToString(plaintext)
}

func Decode(encoded string) []byte {
	bytes, err := xeh.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return bytes
}
