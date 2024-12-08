package http

import (
	"fmt"
	"net"
)

type ResponseWriter struct {
	headerSet bool
	header    Header
	conn      net.Conn
}

func (w *ResponseWriter) Header() Header {
	return w.header
}

func (w *ResponseWriter) WriteHeader(statusCode int) error {
	if w.headerSet {
		panic("Multiple Header writes!")
	}
	w.headerSet = true

	status := fmt.Sprintf("HTTP/1.1 %d %v\r\n", statusCode, StatusText(statusCode))
	_, err := w.conn.Write([]byte(status))
	if err != nil {
		return err
	}
	// write headers
	_, err = w.conn.Write([]byte(w.header.String()))
	if err != nil {
		return err
	}
	_, err = w.conn.Write([]byte("\r\n"))
	return err
}

func (w *ResponseWriter) Write(data []byte) error {
	if !w.headerSet {
		w.WriteHeader(200)
	}
	_, err := w.conn.Write(data)
	return err
}
