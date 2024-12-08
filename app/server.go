package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kjabin/http.go/http"
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
	mux.HandleFunc("POST ", handleFilesPOST)
	log.Fatal(s.ListenAndServe())
}

func handleFilesPOST(w *http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename := os.Args[2] + strings.TrimPrefix(r.Path, "/files/")
	filepath := filename
	err = os.WriteFile(filepath, data, 411)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
	w.Header().Add("Content-Type", "text/plain")

	data := []byte(query)
	encoding, algo, ok := http.ValidEncoding(r.Header.Get("Accept-Encoding"))
	if ok {
		w.Header().Add("Content-Encoding", encoding)
		data = algo(data)
	}
	w.Header().Add("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
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
