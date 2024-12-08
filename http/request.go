package http

import (
	"io"
	"net"
	"strings"
)

type Request struct {
	Method string
	Path   string
	Header Header
	Body   io.Reader
}

func parseRequest(conn net.Conn) (Request, error) {
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return Request{}, err
	}
	sections := strings.Split(string(buf[:n]), "\r\n")
	status := strings.Split(sections[0], " ")

	req := Request{
		Method: status[0],
		Path:   status[1],
		Header: make(Header),
	}

	for _, header := range sections[1 : len(sections)-2] {
		items := strings.Split(header, ": ")
		req.Header.Add(items[0], items[1])
	}
	req.Body = strings.NewReader(sections[len(sections)-1])
	return req, nil
}
