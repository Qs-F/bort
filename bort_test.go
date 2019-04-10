package bort

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"testing"
	"unicode/utf16"
)

func TestBort(t *testing.T) {
	type Case struct {
		Input io.Reader
		Want  bool
	}

	cases := []Case{
		Case{
			Input: bytes.NewReader([]byte("hello, this is the text")),
			Want:  false,
		},
		Case{
			Input: func() *os.File {
				p := filepath.Join("_testdata", "1x1.png")
				f, err := os.Open(p)
				if err != nil {
					t.Fatal("No test data: ", p)
				}
				return f
			}(),
			Want: true,
		},
		// below is the test case for utf16
		// Although this must be successful, the result is false.
		// To use this pkg, make sure always the text file means utf8
		// becuase of golang's regulation
		Case{
			Input: func() io.Reader {
				s := "hello, this is the text"
				buf := make([]byte, 1)
				b := bytes.NewBuffer(buf)
				err := binary.Write(b, binary.LittleEndian, utf16.Encode([]rune(s)))
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
			Want: true,
		},
	}

	for _, c := range cases {
		t.Log("input: ", c.Input)
		b, err := IsBin(c.Input)
		if err != nil {
			t.Fatal(err)
		}
		if b == c.Want {
			t.Log("ok")
		} else {
			t.Fatal("not satisfied")
		}
	}
}
