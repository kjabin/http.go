package http

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Addr string
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", "0.0.0.0"+s.Addr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) error {
	defer conn.Close()
	req, err := parseRequest(conn)
	if err != nil {
		return err
	}

	// connection handlers
	if req.Method == "GET" && req.Path == "/" {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\n")
	} else if req.Method == "GET" && strings.HasPrefix(req.Path, "/echo") {

		query := strings.SplitN(req.Path[1:], "/", 2)[1]
		header := make(Header)
		header.Add("Content-Type", "text/plain")
		header.Add("Content-Length", strconv.Itoa(len(query)))
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n%v\r\n%v", header, query)
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}
	return nil
}

type Header map[string][]string

func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

func (h Header) String() string {
	s := ""
	for k, values := range h {
		for _, v := range values {
			s += fmt.Sprintf("%v: %v\r\n", k, v)
		}
	}
	return s
}

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
