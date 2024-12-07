package http

import (
	"bufio"
	"fmt"
	"net"
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
	req, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	headers := strings.Fields(req)
	if headers[1] == "/" {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\n")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}
	return nil
}
