package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abdirahman04/jama/handlers"
)

func main() {
  fmt.Println("Server listening on port 8080")

  r := http.NewServeMux()

  r.HandleFunc("GET /", handlers.Hello)
  r.HandleFunc("POST /", handlers.Datarer)

  if err := http.ListenAndServe(":8080", r); err != nil {
    log.Fatal(err)
  }
}
