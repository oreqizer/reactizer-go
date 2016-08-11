package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"

	"reactizer-go/server"
	"reactizer-go/modules"
)

func main() {
	server := server.NewServer()
	db, err := sql.Open("postgres", "postgres://oreqizer@localhost/reactizer?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	modules.Register(server, db)
	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
