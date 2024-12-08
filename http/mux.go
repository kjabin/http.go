package http

import "strings"

type HandleFunc func(w *ResponseWriter, r *Request)

type ServeMux struct {
	handlers map[string]HandleFunc
}

func NewServeMux() ServeMux {
	return ServeMux{
		handlers: make(map[string]HandleFunc),
	}
}

func (m *ServeMux) HandleFunc(pattern string, h HandleFunc) {
	if _, ok := m.handlers[pattern]; ok {
		panic("Same pattern registered with ServeMux")
	}
	m.handlers[pattern] = h
}

func (m *ServeMux) Match(r Request) HandleFunc {
	for pattern, f := range m.handlers {
		p := strings.Split(pattern, " ")
		method, pattern := p[0], p[1]
		if method == r.Method && strings.HasPrefix(r.Path, pattern) {
			return f
		}
	}
	return handleDefault
}

func handleDefault(w *ResponseWriter, r *Request) {
	w.WriteHeader(405)
}
