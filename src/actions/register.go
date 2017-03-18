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

func Register(cfg *ini.File, data RegisterData) (error) {
	var db *sql.DB
	var err error

	dbtype := cfg.Section("general").Key("dbstore").String()
	switch dbtype {
	case "sqlite":
		location := cfg.Section("sqlite").Key("location").String()
		db, err = dbstore.OpenSQLite(location)
	case "postgresql":
		dbhost := cfg.Section("postgresql").Key("host").String()
		dbport, err := cfg.Section("postgresql").Key("port").Int()
		if err != nil {
			return errors.New("Invalid PostgreSQL port: " + err.Error())
		}
		dbuser := cfg.Section("postgresql").Key("user").String()
		dbpassword := cfg.Section("postgresql").Key("password").String()
		dbdatabase := cfg.Section("postgresql").Key("database").String()
		db, err = dbstore.OpenPostgres(dbhost, dbport, dbuser, dbpassword, dbdatabase)
	default:
		return errors.New("Database not supported: " + dbtype)
	}

	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
