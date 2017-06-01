/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	"strings"
)

func tables() []string {
	return []string{
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
			"hostname VARCHAR(30),",
			"assigned TIMESTAMP)",
		}, " "),
	}
}

func InsertSection(db *sql.DB, section, network, netmask string) error {
	var err error

	query := "INSERT INTO sections(section, network, netmask) VALUES ($1, $2, $3)"

	tx, _ := db.Begin()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(section, network, netmask)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	return nil
}

func IsSectionExists(db *sql.DB, section string) (bool, error) {
	var counted int

	query := "SELECT count(*) FROM sections WHERE section = $1"
	if err := db.QueryRow(query, section).Scan(&counted); err != nil {
		return false, err
	}

	if counted > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func RemoveSection(db *sql.DB, section string) error {
	query := "DELETE FROM sections WHERE section = $1"
	// TODO: remove data from children table

	_, err := db.Exec(query, section)

	return err
}
