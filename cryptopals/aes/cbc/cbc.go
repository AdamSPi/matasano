package cbc

import (
	"bytes"
	"crypto/aes"
	"matasano/cryptopals/aes/ecb"
	"matasano/cryptopals/util"
	"matasano/cryptopals/xor"
)

func Encrypt(plaintext, password string) ([]byte, error) {
	blockCipher, err := aes.NewCipher([]byte(password))
	if err != nil {
		return nil, err
	}
	blockSize := blockCipher.BlockSize()
	paddedText, err := util.Pad([]byte(plaintext), blockSize)
	if err != nil {
		return nil, err
	}

	chunks := util.Chunkify(paddedText, blockSize)
	var buffer bytes.Buffer
	initial := bytes.Repeat([]byte{0}, blockSize)
	vector := initial

	for _, chunk := range chunks {
		xorChunk := xor.EncryptBytes(chunk, vector)
		cipherChunk, err := ecb.EncryptBytes(xorChunk, password)
		if err != nil {
			return nil, err
		}
		buffer.Write(cipherChunk)
		vector = cipherChunk
	}
	return buffer.Bytes(), nil
}

func Decrypt(ciphertext []byte, password string) (string, error) {
	blockCipher, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	blockSize := blockCipher.BlockSize()

	chunks := util.Chunkify(ciphertext, blockSize)
	var plaintext string
	initial := bytes.Repeat([]byte{0}, blockSize)
	vector := initial

	for _, chunk := range chunks {
		cipherChunk, err := ecb.DecryptBytes(chunk, password)
		if err != nil {
			return "", err
		}
		xorChunk := xor.EncryptBytes(cipherChunk, vector)
		plaintext = plaintext + string(xorChunk)
		vector = chunk
	}

	return plaintext, nil
}
