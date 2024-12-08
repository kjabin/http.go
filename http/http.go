package http

import (
	"net"
)

type Server struct {
	Addr    string
	Handler ServeMux
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
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	r, err := parseRequest(conn)
	if err != nil {
		return
	}

	f := s.Handler.Match(r)
	w := ResponseWriter{
		headerSet: false,
		header:    make(Header),
		conn:      conn,
	}
	f(&w, &r)
}
