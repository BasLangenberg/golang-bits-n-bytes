package persistence

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	// sqlite driver will be called using sqlx helper library
	_ "github.com/mattn/go-sqlite3"
)

// Init initialized db if not existing, does schema creation and returns a pointer to the sqlx.DB instance
func Init() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", os.Getenv("PWD")+"/db.sqlite")
	if err != nil {
		return nil, fmt.Errorf("unable to open sqlite database: %+v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to sqlite database")
	}

	return db, nil
}
