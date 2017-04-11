/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"strings"
	"database/sql"
	"errors"
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

func InsertSection(db *sql.DB, dbtype string, section, network, netmask string) error {
	var err error

	query := "INSERT INTO sections(section, network, netmask) VALUES"

	switch dbtype {
	case "sqlite":
		query = query + "(?, ?, ?)"
	case "postgresql":
		query = query + "($1, $2, $3)"
	default:
		return errors.New("No database engine supported")
	}

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
