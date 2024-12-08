package http

import (
	"bytes"
	"compress/gzip"
	"strings"
)

type Algorithm func([]byte) []byte

var algorithms map[string]Algorithm = map[string]Algorithm{
	"gzip": zip,
}

func ValidEncoding(encoding string) (string, Algorithm, bool) {
	encodings := strings.Split(encoding, ", ")
	for _, encoding := range encodings {
		algorithm, ok := algorithms[encoding]
		if ok {
			return encoding, algorithm, ok
		}
	}
	return "", nil, false
}

func zip(data []byte) []byte {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		panic(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
