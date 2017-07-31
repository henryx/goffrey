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
	"strconv"
)

func isMySQLExists(db *sql.DB) bool {
	// TODO: implement check

	return true
}

func OpenMySQL(host string, port int, user string, password string, database string) (*sql.DB, error) {

	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if !isMySQLExists(db) {

	}

	return db, nil
}
