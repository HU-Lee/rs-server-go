package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func DatabaseInit() {
	db_url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		panic(err)
	}

	PostgresDB = db
	PostgresDB.SetMaxIdleConns(10)
}
