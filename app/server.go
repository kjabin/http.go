package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

func main() {
	mux := http.NewServeMux()
	s := http.Server{
		Addr:    ":4221",
		Handler: mux,
	}
	mux.HandleFunc("GET /echo", handleEcho)
	mux.HandleFunc("GET /user-agent", handleUserAgent)
	mux.HandleFunc("GET /", handleIndex)
	log.Fatal(s.ListenAndServe())
}

func handleIndex(w *http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func handleEcho(w *http.ResponseWriter, r *http.Request) {
	query := strings.SplitN(r.Path[1:], "/", 2)[1]
	w.Header().Add("Content-Length", strconv.Itoa(len(query)))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte(query))
}

func handleUserAgent(w *http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	w.Header().Add("Content-Length", strconv.Itoa(len(userAgent)))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte(userAgent))
}

func handleDefault(w *http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}
