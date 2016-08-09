package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("hello world")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
