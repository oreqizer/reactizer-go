package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"

	"reactizer-go/server"
	"reactizer-go/modules"
	"reactizer-go/i18n"
	"reactizer-go/config"
)

func main() {
	i18n.LoadTranslations(config.Locales)

	server := server.NewServer()
	db, err := sql.Open("postgres", config.DBurl + "?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	modules.Register(server, db)
	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
