/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func isMySQLExists(db *sql.DB, dbname string) bool {
	var counted int

	query := "SELECT Count(*) FROM information_schema.tables " +
		"WHERE table_schema = ? " +
		"AND table_name IN ('status', 'attrs', 'acls')"

	err := db.QueryRow(query, dbname).Scan(&counted)
	if err != nil {
		log.Fatal("Schema 1: Error in check database structure: " + err.Error())
	}

	if counted > 0 {
		return true
	} else {
		return false
	}

	return true
}

func OpenMySQL(host string, port int, user string, password string, database string) (*sql.DB, error) {

	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if !isMySQLExists(db, database) {

	}

	return db, nil
}
