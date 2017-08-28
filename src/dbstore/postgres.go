/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

func isPgDbExists(db *sql.DB) bool {
	var counted int

	query := "SELECT Count(*) FROM information_schema.tables " +
		"WHERE table_schema = 'public' " +
		"AND table_name in ('sections, 'addresses')"

	err := db.QueryRow(query).Scan(&counted)
	if err != nil {
		log.Fatal("Schema 1: Error in check database structure: " + err.Error())
	}

	if counted > 0 {
		return true
	} else {
		return false
	}
}

func OpenPostgres(host string, port int, user string, password string, database string) (*sql.DB, error) {
	dsn := strings.Join([]string{
		"user=" + user,
		"password=" + password,
		"host=" + host,
		"port=" + strconv.Itoa(port),
		"dbname=" + database,
		"sslmode=disable",
	}, " ")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if !isPgDbExists(db) {
		createDb(db)
	}

	return db, nil
}
