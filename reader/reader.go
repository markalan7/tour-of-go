package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read (b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i,v := range b {
		switch {
		case v >= 'a' && v <= 'm':
			b[i] += 13
		case v >= 'A' && v <= 'M':
			b[i] += 13
		case v >= 'n' && v <= 'z':
			b[i] -= 13
		case v >= 'N' && v <= 'Z':
			b[i] -= 13
		default:

		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
