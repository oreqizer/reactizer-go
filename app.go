package main

import (
	"reactizer-go/server"
	"reactizer-go/modules"
	"log"
	"net/http"
)

func main() {
	server := server.NewServer()
	modules.MountRoutes(server)
	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
