package main

import (
	"log"
	"os"
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
	mux.HandleFunc("GET ", handleDefault)
	mux.HandleFunc("GET /files", handleFiles)
	log.Fatal(s.ListenAndServe())
}

func handleFiles(w *http.ResponseWriter, r *http.Request) {
	filename := strings.SplitN(r.Path[1:], "/", 2)[1]
	if _, err := os.Stat(filename); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
}

func handleIndex(w *http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleEcho(w *http.ResponseWriter, r *http.Request) {
	query := strings.SplitN(r.Path[1:], "/", 2)[1]
	w.Header().Add("Content-Length", strconv.Itoa(len(query)))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(query))
}

func handleUserAgent(w *http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	w.Header().Add("Content-Length", strconv.Itoa(len(userAgent)))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userAgent))
}

func handleDefault(w *http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
