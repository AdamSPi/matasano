package xor

import (
	"matasano/cryptopals/b64"
	"matasano/cryptopals/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestXorRoundTrip(t *testing.T) {
	cases := []struct {
		name      string
		plaintext string
		password  string
	}{
		{
			"sample test",
			"Aphex Twin - Rare Ambient Works mixtape",
			"syrobonkus",
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				ciphertext := Encrypt(c.plaintext, c.password)
				if diff := cmp.Diff(Decrypt(ciphertext, c.password), c.plaintext); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestXorRoundTripFile(t *testing.T) {
	cases := []struct {
		name      string
		filepath  string
		plaintext string
		password  string
	}{
		{
			"challenge 6 file test",
			"../../test/challenge6.txt",
			"../../test/play_that_funky_music.txt",
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
				plaintext := Decrypt(decodedText, c.password)
				if diff := cmp.Diff(plaintext, string(util.ReadFile(c.plaintext))); diff != "" {
					t.Error(diff)
				}
				encryptedText := Encrypt(plaintext, c.password)
				encodedText := []byte(b64.Encode(encryptedText))
				if diff := cmp.Diff(encodedText, ciphertext); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}
