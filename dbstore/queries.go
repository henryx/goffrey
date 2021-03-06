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
)

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

func IsHostExists(db *sql.DB, section, hostname string) (bool, error) {
	var counted int

	query := "SELECT count(*) FROM addresses WHERE section = ? and hostname = ?"
	if err := db.QueryRow(query, section, hostname).Scan(&counted); err != nil {
		return false, err
	}

	if counted > 0 {
		return true, nil
	} else {
		return false, nil
	}
}


func RetrieveFreeIP(db *sql.DB, section string) (string, error) {
	var result string
	query := "SELECT address FROM addresses WHERE section = ? AND hostname IS NULL ORDER BY address LIMIT 1"

	if err := db.QueryRow(query, section).Scan(&result); err != nil {
		return "", err
	}

	return result, nil
}