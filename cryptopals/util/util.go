package util

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("could not read file")
	}
	return data
}

func ScanFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	data := make([]byte, 0, 1024)

	for scan.Scan() {
		input := []byte(scan.Text())
		data = append(data, input...)
	}
	return data
}

func Chunkify(text []byte, size int) [][]byte {
	chunks := make([][]byte, 0, len(text)/size+1)
	for size < len(text) {
		chunks = append(chunks, text[:size])
		text = text[size:]
	}
	return append(chunks, text)
}

func Blockify(text []byte, size int) [][]byte {
	if len(text) == 0 {
		return nil
	}
	blocks := make([][]byte, 0, size)
	for offset := 0; offset < size; offset++ {
		var block []byte
		var chunk int
		blockIndex := offset + chunk*size
		for blockIndex < len(text) {
			block = append(block, text[blockIndex])
			chunk++
			blockIndex = offset + chunk*size
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func Pad(text []byte, size int) ([]byte, error) {
	if len(text)%size == 0 {
		return text, nil
	}
	if size <= 1 || size > 255 {
		return text, errors.New("invalid block size")
	}
	n := size - len(text)%size
	padding := []byte{byte(n)}
	return append(text, bytes.Repeat(padding, n)...), nil
}

func Unpad(text []byte, size int) ([]byte, error) {
	if len(text)%size != 0 {
		return text, errors.New("invalid padding")
	}
	n := int(text[len(text)-1])
	padding := bytes.Repeat([]byte{byte(n)}, n)
	if !bytes.Equal(padding, text[len(text)-n:]) {
		return nil, errors.New("incorrect padding")
	}
	return text[:len(text)-n], nil
}
