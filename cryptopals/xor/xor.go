package xor

func Encrypt(plaintext, password string) []byte {
	text, keyword := []byte(plaintext), []byte(password)
	return vigenereCipher(text, keyword)
}

func Decrypt(ciphertext []byte, password string) string {
	keyword := []byte(password)
	return string(vigenereCipher(ciphertext, keyword))
}

func EncryptBytes(plaintext, keyword []byte) []byte {
	return vigenereCipher(plaintext, keyword)
}

func vigenereCipher(text, keyword []byte) []byte {
	ciphertext := make([]byte, len(text))
	for index, character := range text {
		ciphertext[index] = character ^ keyword[index%len(keyword)]
	}
	return ciphertext
}

func caesarCipher(text []byte, key byte) []byte {
	result := make([]byte, len(text))
	for index, character := range text {
		result[index] = character ^ key
	}
	return result
}
