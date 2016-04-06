package main

import (
	"bytes"
	"io"
	"os"
)

func Asset(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	return buf.Bytes(), err
}
