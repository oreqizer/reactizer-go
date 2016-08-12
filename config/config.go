package config

import "flag"

var (
	Locales = []string{"en-US", "sk"}
	DBurl = *flag.String("db", "postgres://oreqizer@localhost/reactizer", "Database URL")
	Port = *flag.Int("port", 8080, "Server port")
)

const (
	DefaultLanguage = "en-US"
)
