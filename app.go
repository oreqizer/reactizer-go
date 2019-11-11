package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"reactizer-go/api"
	"reactizer-go/config"
	"reactizer-go/i18n"

	"github.com/kataras/iris/v12"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	i18n.LoadTranslations(config.Locales)

	app := iris.New()
	db, err := sql.Open("postgres", config.DBurl+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	api.Register(app, db)

	port := fmt.Sprintf(":%d", config.Port)
	app.Run(iris.Addr(port))
}
