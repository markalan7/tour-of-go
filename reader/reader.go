package main

import "code.google.com/p/go-tour/reader"
    
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
    for i := range b {
        b[i] = byte('A')
    }
    return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
