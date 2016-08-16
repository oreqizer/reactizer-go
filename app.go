package main

import (
	"log"
	"fmt"
	"database/sql"

	"github.com/kataras/iris"
	_ "github.com/lib/pq"

	"reactizer-go/api"
	"reactizer-go/i18n"
	"reactizer-go/config"
	"flag"
)

func main() {
	flag.Parse()
	i18n.LoadTranslations(config.Locales)

	app := iris.New()
	db, err := sql.Open("postgres", config.DBurl + "?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	api.Register(app, db)

	port := fmt.Sprintf(":%d", config.Port)
	app.Listen(port)
}
