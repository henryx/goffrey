/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

import (
	"database/sql"
	"dbstore"
	"errors"
	"github.com/go-ini/ini"
	"github.com/op/go-logging"
)

func openDb(log *logging.Logger, cfg *ini.File) (*sql.DB, error) {
	var db *sql.DB
	var err error

	dbtype := cfg.Section("general").Key("database").String()
	switch dbtype {
	case "sqlite":
		db, err = openSqlite(cfg.Section("sqlite").Key("location").String())
	case "postgres":
		sect := cfg.Section("postgres")
		db, err = openPg(sect)
	default:
		log.Debug("Database specified: " + dbtype)
		return nil, errors.New("Database not supported")
	}

	return db, err
}

func openSqlite(location string) (*sql.DB, error) {
	db, err := dbstore.OpenSQLite(location)
	return db, err
}

func openPg(sect *ini.Section) (*sql.DB, error) {
	dbhost := sect.Key("host").String()
	dbport, err := sect.Key("port").Int()
	if err != nil {
		return nil, errors.New("Invalid PostgreSQL port: " + err.Error())
	}

	dbuser := sect.Key("user").String()
	dbpassword := sect.Key("password").String()
	dbdatabase := sect.Key("database").String()

	db, err := dbstore.OpenPostgres(dbhost, dbport, dbuser, dbpassword, dbdatabase)

	return db, err
}
