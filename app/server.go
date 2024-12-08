package main

import (
	"fmt"
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
	mux.HandleFunc("GET /files", handleFiles)
	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET ", handleDefault)
	log.Fatal(s.ListenAndServe())
}

func handleFiles(w *http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.Path, "/files/")
	filepath := os.Args[2] + filename
	if _, err := os.Stat(filepath); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleIndex(w *http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleEcho(w *http.ResponseWriter, r *http.Request) {
	query := strings.TrimPrefix(r.Path, "/echo/")
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
