package main

import (
	"log"
	"net/http"
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"

	"reactizer-go/server"
	"reactizer-go/modules"
	"reactizer-go/i18n"
	"reactizer-go/config"
)

func main() {
	i18n.LoadTranslations(config.Locales)

	box := server.NewServeBox()
	db, err := sql.Open("postgres", config.DBurl + "?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	modules.Register(box, db)

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, box.Mux))
}
