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
	"errors"
	"dbstore"
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

	err = dbstore.InsertSection(db, dbtype, data.Name, data.Network, data.Netmask)
	if err != nil {
		return err
	} else {
		return nil
	}
}
