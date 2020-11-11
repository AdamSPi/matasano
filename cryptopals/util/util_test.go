package util

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunkify(t *testing.T) {
	cases := []struct {
		name     string
		text     []byte
		size     int
		expected [][]byte
	}{
		{
			"short string",
			[]byte("abc"),
			3,
			[][]byte{[]byte("abc")},
		},
		{
			"empty string",
			[]byte(""),
			1,
			[][]byte{[]byte("")},
		},
		{
			"long string",
			[]byte("g19403g134ng341"),
			2,
			[][]byte{
				[]byte("g1"), []byte("94"), []byte("03"), []byte("g1"),
				[]byte("34"), []byte("ng"), []byte("34"), []byte("1"),
			},
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				if diff := cmp.Diff(Chunkify(c.text, c.size), c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestBlockify(t *testing.T) {
	cases := []struct {
		name     string
		text     []byte
		size     int
		expected [][]byte
	}{

		{
			"long string",
			[]byte("g19403g134ng341"),
			2,
			[][]byte{
				[]byte("g90g3n31"), []byte("14314g4"),
			},
		},
		{
			"empty string",
			[]byte(""),
			1,
			nil,
		},
		{
			"short string",
			[]byte("aaabbbccc"),
			3,
			[][]byte{[]byte("abc"), []byte("abc"), []byte("abc")},
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				out := Blockify(c.text, c.size)
				if diff := cmp.Diff(out, c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestPadding(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		size     int
		expected []byte
	}{

		{
			"challenge 9 test",
			"YELLOW SUBMARINE",
			20,
			append(
				[]byte("YELLOW SUBMARINE"), bytes.Repeat([]byte{4}, 4)...,
			),
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				out, _ := Pad([]byte(c.text), c.size)
				if diff := cmp.Diff(out, c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestUnpadding(t *testing.T) {
	cases := []struct {
		name     string
		text     []byte
		size     int
		expected []byte
	}{

		{
			"unpad test",
			append(
				[]byte("YELLOW SUBMARINE"), bytes.Repeat([]byte{4}, 4)...,
			),
			20,
			[]byte("YELLOW SUBMARINE"),
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				out, _ := Unpad(c.text, c.size)
				if diff := cmp.Diff(out, c.expected); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}

func TestPaddingFile(t *testing.T) {
	cases := []struct {
		name           string
		size           int
		filepath       string
		paddedFilepath string
	}{
		{
			"pad file test",
			16,
			"../../test/play_that_funky_music.txt",
			"../../test/play_that_funky_music_padded.txt",
		},
	}

	for _, _c := range cases {
		c := _c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				unpadded := ReadFile(c.filepath)
				padded, _ := Pad(unpadded, c.size)
				if diff := cmp.Diff(padded, ReadFile(c.paddedFilepath)); diff != "" {
					t.Error(diff)
				}
				undo, _ := Unpad(padded, c.size)
				if diff := cmp.Diff(undo, unpadded); diff != "" {
					t.Error(diff)
				}
			},
		)
	}
}
