package ecb

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(plaintext, password string) ([]byte, error) {
	blockCipher, err := aes.NewCipher([]byte(password))
	if err != nil {
		return nil, err
	}
	encrypter := newEncrypter(blockCipher)
	return encrypter.crypt([]byte(plaintext)).Bytes(), nil
}

func Decrypt(ciphertext []byte, password string) (string, error) {
	blockCipher, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	decrypter := newDecrypter(blockCipher)
	return decrypter.crypt(ciphertext).String(), nil
}

func newEncrypter(cipher cipher.Block) codebook {
	return codebook{
		cipher,
		cipher.Encrypt,
		cipher.BlockSize(),
	}
}

func newDecrypter(cipher cipher.Block) codebook {
	return codebook{
		cipher,
		cipher.Decrypt,
		cipher.BlockSize(),
	}
}

type codebook struct {
	mode    cipher.Block
	crypter func([]byte, []byte)
	size    int
}

func (b *codebook) crypt(ciphertext []byte) *bytes.Buffer {
	var buffer bytes.Buffer
	for {
		block := make([]byte, b.size)
		b.crypter(block, ciphertext[:b.size])
		buffer.Write(block)
		if ciphertext = ciphertext[b.size:]; len(ciphertext) == 0 {
			break
		}
	}
	return &buffer
}
