package http

import (
	"strings"
)

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
	bestMatch, bestFunc := -1, handleDefault
	for pattern, f := range m.handlers {
		p := strings.Split(pattern, " ")
		method, pattern := p[0], p[1]
		if method != r.Method {
			continue
		}
		if !strings.HasPrefix(r.Path, pattern) {
			continue
		}
		if strings.HasSuffix(pattern, "/") && pattern != r.Path {
			continue
		}
		if len(pattern) >= bestMatch {
			bestMatch, bestFunc = len(pattern), f
		}
	}
	return bestFunc
}

func handleDefault(w *ResponseWriter, r *Request) {
	w.WriteHeader(405)
}
