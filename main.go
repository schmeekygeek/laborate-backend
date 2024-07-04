package main

import (
	"log"
	"net/http"
)

func main() {
  server := Init()
  http.HandleFunc("/", server.Serve)
  log.Println("Starting server on port :8080")
  http.ListenAndServe(":8080", nil)
}
