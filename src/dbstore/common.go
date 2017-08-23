/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	"ip"
	"log"
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

func InsertSection(db *sql.DB, section, network, netmask string) error {
	var err error

	queries := []string{
		"INSERT INTO sections(section, network, netmask) VALUES ($1, $2, $3)",
		"INSERT INTO addresses(section, address) VALUES($1, $2)",
	}

	tx, _ := db.Begin()

	stmt, err := tx.Prepare(queries[0])
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(section, network, netmask)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err := tx.Prepare(queries[1])
	if err != nil {
		return err
	}
	defer stmt2.Close()

	mask, err := ip.ToCidr(netmask)
	if err != nil {
		tx.Rollback()
		return err
	}

	ips, err := ip.Range(network, mask)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, ipaddr := range ips {
		_, err = stmt2.Exec(section, ipaddr)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
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
	queries := []string{
		"DELETE FROM addresses WHERE section = $1",
		"DELETE FROM sections WHERE section = $1",
	}

	tx, _ := db.Begin()

	for _, query := range queries {
		stmt, err := tx.Prepare(query)

		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		_, err = tx.Exec(query, section)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
