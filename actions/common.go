/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

import (
	"database/sql"
	"errors"
	"github.com/go-ini/ini"
	"goffrey/dbstore"
)

func openDb(cfg *ini.File) (*sql.DB, error) {
	dbhost := cfg.Section("database").Key("host").String()
	dbport, err := cfg.Section("database").Key("port").Int()
	if err != nil {
		return nil, errors.New("Invalid MySQL port: " + err.Error())
	}
	dbuser := cfg.Section("database").Key("user").String()
	dbpassword := cfg.Section("database").Key("password").String()
	dbdatabase := cfg.Section("database").Key("name").String()

	db, err := dbstore.OpenMySQL(dbhost, dbport, dbuser, dbpassword, dbdatabase)

	return db, err
}

func checkHost(db *sql.DB, section, host string) bool {
	var hostexists bool
	var err error

	hostexists, err = dbstore.IsHostExists(db, section, host)
	if err != nil {
		return false
	}

	if hostexists {
		return true
	} else {
		return false
	}
}
