package test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(t *testing.T) *sql.DB {
	t.Helper()

	conn := os.Getenv("DB_HOST")
	if conn == "" {
		t.Skip("DB_HOST not set")
	}

	db, err := sql.Open("pgx", conn)
	if err != nil {
		t.Fatal(err)
	}

	return db
}
