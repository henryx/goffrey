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

func isMySQLExists(db *sql.DB, dbname string) bool {
	var counted int

	query := "SELECT Count(*) FROM information_schema.tables " +
		"WHERE table_schema = ? " +
		"AND table_name IN ('sections', 'addresses')"

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
		createDb(db)
	}

	return db, nil
}

func IsSectionExists(db *sql.DB, section string) (bool, error) {
	var counted int

	query := "SELECT count(*) FROM sections WHERE section = ?"
	if err := db.QueryRow(query, section).Scan(&counted); err != nil {
		return false, err
	}

	if counted > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
