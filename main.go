package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err)
  }
  log.Print("Listening on 8080")
}
