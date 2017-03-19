/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

import (
	"github.com/go-ini/ini"
	"database/sql"
	"dbstore"
	"errors"
)

type RegisterData struct {
	Enable  bool
	Name    string `names:"-n, --name" usage:"Set the name of the network"`
	Network string `names:"-N, --network" usage:"Set the network addresses"`
	Netmask string `names:"-M, --netmask" usage:"Set the network mask"`
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

func Register(cfg *ini.File, data RegisterData) (error) {
	var db *sql.DB
	var err error

	dbtype := cfg.Section("general").Key("dbstore").String()
	switch dbtype {
	case "sqlite":
		db, err = openSqlite(cfg.Section("sqlite").Key("location").String())
	case "postgresql":
		sect := cfg.Section("postgresql")
		db, err = openPg(sect)
	default:
		return errors.New("Database not supported: " + dbtype)
	}

	if err != nil {
		return err
	}
	defer db.Close()

	//TODO: continue it
	return nil
}
