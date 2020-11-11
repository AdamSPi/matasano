package ecb

import (
	"matasano/cryptopals/b64"
	"matasano/cryptopals/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestElectronicCodebookRoundTripFile(t *testing.T) {
	cases := []struct {
		name      string
		filepath  string
		plaintext string
		password  string
	}{
		{
			"challenge 7 test",
			"../../../test/challenge7.txt",
			"../../../test/play_that_funky_music_padded.txt",
			"YELLOW SUBMARINE",
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
				plaintext, _ := Decrypt(decodedText, c.password)
				if diff := cmp.Diff(plaintext, string(util.ReadFile(c.plaintext))); diff != "" {
					t.Error(diff)
				}
				encryptedText, _ := Encrypt(plaintext, c.password)
				encodedText := []byte(b64.Encode(encryptedText))
				if diff := cmp.Diff(string(encodedText), string(ciphertext)); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}
