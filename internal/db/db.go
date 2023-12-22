package db

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDb(dbFile string) *sql.DB {

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	seedDb(db)

	return db
}
