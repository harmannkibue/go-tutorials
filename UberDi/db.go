package UberDi

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}
