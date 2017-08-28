/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	_ "github.com/gwenn/gosqlite"
	"log"
)

func isSQLiteDbExist(db *sql.DB) bool {
	var counted int
	query := "SELECT count(*) FROM sqlite_master"

	if err := db.QueryRow(query).Scan(&counted); err != nil {
		log.Fatal("Schema 1: Failed to check schema database: " + err.Error())
	}

	if counted > 0 {
		return true
	} else {
		return false
	}
}

func OpenSQLite(location string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}

	if !isSQLiteDbExist(db) {
		createDb(db)
	}
	return db, nil
}
