package db

import (
	"database/sql"
	"fmt"
	"os"
)

func seedDb(db *sql.DB) {

	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='items')")
	var tableExists bool
	err := row.Scan(&tableExists)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if tableExists {
		return
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL, 
			created_at DATE DEFAULT CURRENT_TIMESTAMP NOT NULL
		)`,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(
		`INSERT INTO items (name) VALUES ('Item 1'), ('Item 2')`,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
