package main

import (
	"fmt"
	"matasano/cryptopals/aes/ecb"
	"matasano/cryptopals/b64"
	"matasano/cryptopals/util"
	"os"
)

func main() {
	ciphertext := util.ScanFile(os.Args[1])
	encodedText := b64.Decode(string(ciphertext))
	password := os.Args[2]
	plaintext, _ := ecb.Decrypt(encodedText, password)
	fmt.Print(plaintext)
}
