/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

func createDb(db *sql.DB) {
	var tx *sql.Tx
	var tables = []string{
		strings.Join([]string{
			"CREATE TABLE sections(",
			"section VARCHAR(30),",
			"network VARCHAR(15),",
			"netmask VARCHAR(15),",
			"PRIMARY KEY(section))",
		}, " "),
		strings.Join([]string{
			"CREATE TABLE addresses(",
			"section VARCHAR(30),",
			"address VARCHAR(15),",
			"hostname VARCHAR(255),",
			"assigned TIMESTAMP)",
		}, " "),
	}

	tx, _ = db.Begin()
	for _, query := range tables {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			break
		}
	}
	tx.Commit()
}

func OpenMySQL(host string, port int, user string, password string, database string) (*sql.DB, error) {

	dsn := user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if !isMySQLExists(db, database) {
		createDb(db)
	}

	return db, nil
}
