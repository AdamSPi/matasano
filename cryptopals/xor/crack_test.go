package xor

import (
	"matasano/cryptopals/b64"
	"matasano/cryptopals/hex"
	"matasano/cryptopals/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHammingDistance(t *testing.T) {
	t.Run(
		"hamming distance example test",
		func(t *testing.T) {
			if d := distance([]byte("this is a test"), []byte("wokka wokka!!!")); d != 37 {
				t.Errorf("expected distance mismatch: expected 37, got %d", d)
			}
		},
	)
}

func TestKeySizeScoreHeap(t *testing.T) {
	cases := []struct {
		name      string
		plaintext string
		key       string
		head      int
		expected  keySizeHeap
	}{
		{
			"trusted signal test",
			"fuse fuel for falling flocks",
			"few",
			5,
			keySizeHeap{
				{
					12,
					2.1666666666666665,
				},
				{
					11,
					2.1818181818181817,
				},
				{
					9,
					2.2222222222222223,
				},
				{
					6,
					2.444444444444444,
				},
				{
					3,
					2.5416666666666665,
				},
			},
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				encrypted := Encrypt(c.plaintext, c.key)
				out := guessKeySizes(encrypted, c.head)
				if diff := cmp.Diff(out, c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestBruteforceAttack(t *testing.T) {
	cases := []struct {
		name       string
		ciphertext string
		expected   byte
	}{
		{
			"challenge 3 test",
			"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			byte(88),
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				hexDecode := hex.Decode(c.ciphertext)
				out := crackKey(hexDecode)
				if diff := cmp.Diff(out, c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestCrack(t *testing.T) {
	cases := []struct {
		name      string
		plaintext string
		expected  string
	}{
		{
			"manual test",
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			"I'm pickle rick!",
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				ciphertext := Encrypt(c.plaintext, c.expected)
				out := Crack(ciphertext)
				if diff := cmp.Diff(string(out[0]), c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestCrackFile(t *testing.T) {
	cases := []struct {
		name     string
		filepath string
		expected string
	}{
		{
			"challenge 6 test",
			"../../test/challenge6.txt",
			"Terminator X: Bring the noise",
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				ciphertext := util.ScanFile(c.filepath)
				decodedText := b64.Decode(string(ciphertext))
				out := Crack(decodedText)
				if diff := cmp.Diff(string(out[0]), c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}
