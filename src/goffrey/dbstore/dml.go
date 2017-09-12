/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	"goffrey/ip"
)

func InsertSection(db *sql.DB, section, network, netmask string) error {
	var err error

	queries := []string{
		"INSERT INTO sections(section, network, netmask) VALUES (?, ?, ?)",
		"INSERT INTO addresses(section, address) VALUES(?, ?)",
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

func RemoveSection(db *sql.DB, section string) error {
	queries := []string{
		"DELETE FROM addresses WHERE section = ?",
		"DELETE FROM sections WHERE section = ?",
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

func AssignHost(db *sql.DB, section, hostname, ipaddr string) error {
	update := "UPDATE addresses SET hostname = ?, assigned = CURRENT_TIMESTAMP WHERE section = ? AND address = ?"

	tx, _ := db.Begin()

	stmt, err := tx.Prepare(update)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = tx.Exec(update, hostname, section, ipaddr)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetHost(db *sql.DB, section, ipaddr string) (string, error) {
	var count int
	var host string
	var err error

	counted := "SELECT count(hostname) FROM addresses WHERE section = ? AND address = ?"
	query := "SELECT hostname FROM addresses WHERE section = ? AND address = ?"

	err = db.QueryRow(counted, section, ipaddr).Scan(&count)
	if err != nil {
		return "", err
	}

	if count == 0 {
		return "", nil
	}

	err = db.QueryRow(query, section, ipaddr).Scan(&host)
	if err != nil {
		return "", err
	}

	return host, nil
}

func GetIP(db *sql.DB, section, hostname string) (string, error) {
	var count int
	var ipaddr string
	var err error

	counted := "SELECT count(address) FROM addresses WHERE section = ? AND hostname = ?"
	query := "SELECT address FROM addresses WHERE section = ? AND hostname = ?"

	err = db.QueryRow(counted, section, hostname).Scan(&count)
	if err != nil {
		return "", err
	}

	if count == 0 {
		return "", nil
	}

	err = db.QueryRow(query, section, hostname).Scan(&ipaddr)
	if err != nil {
		return "", err
	}

	return ipaddr, nil
}
