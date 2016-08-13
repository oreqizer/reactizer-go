package config

import "flag"

var (
	Locales = []string{"en-US", "sk"}

	// flaggable:
	DBurl = *flag.String("db", "postgres://oreqizer@localhost/reactizer", "Database URL")
	Port = *flag.Int("port", 8080, "Server port")
	Secret = *flag.String("secret", "wowMuchSecure1337", "Secret")
)

const (
	DefaultLanguage = "en-US"
)
