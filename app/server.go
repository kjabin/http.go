package main

import (
	"log"

	"github.com/codecrafters-io/http-server-starter-go/http"
)

func main() {
	s := http.Server{Addr: ":4221"}
	log.Fatal(s.ListenAndServe())
}
