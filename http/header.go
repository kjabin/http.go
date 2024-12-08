package http

import "fmt"

type Header map[string][]string

func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

func (h Header) Get(key string) string {
	if _, ok := h[key]; !ok || len(h[key]) == 0 {
		return ""
	}
	return h[key][0]
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
