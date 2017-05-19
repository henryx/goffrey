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
	"encoding/json"
	"errors"
	"github.com/go-ini/ini"
)

type RegisterData struct {
	Enable  bool
	Name    string `names:"-n, --name" usage:"Set the name of the network"`
	Network string `names:"-N, --network" usage:"Set the network addresses"`
	Netmask string `names:"-M, --netmask" usage:"Set the network mask"`
}

func Register(cfg *ini.File, data RegisterData) error {
	var db *sql.DB
	var err error

	if data.Name == "" {
		jsonerr, _ := json.Marshal(&ActionError{
			Action:  "register",
			Code:    1,
			Message: "No section name passed",
		})
		return errors.New(string(jsonerr))
	}

	dbtype := cfg.Section("general").Key("database").String()
	switch dbtype {
	case "sqlite":
		db, err = openSqlite(cfg.Section("sqlite").Key("location").String())
	case "postgres":
		sect := cfg.Section("postgres")
		db, err = openPg(sect)
	default:
		jsonerr, _ := json.Marshal(&ActionError{
			Action:  "register",
			Code:    2,
			Message: "Database not supported:" + dbtype,
		})
		return errors.New(string(jsonerr))
	}

	if err != nil {
		return err
	}
	defer db.Close()

	exists := dbstore.IsSectionExists(db, data.Name)
	if exists {
		jsonerr, _ := json.Marshal(&ActionError{
			Action:  "register",
			Code:    2,
			Message: "Section " + data.Name + " already exists",
		})
		return errors.New(string(jsonerr))
	}

	err = dbstore.InsertSection(db, data.Name, data.Network, data.Netmask)
	if err != nil {
		jsonerr, _ := json.Marshal(&ActionError{
			Action:  "register",
			Code:    3,
			Message: err.Error(),
		})
		return errors.New(string(jsonerr))
	} else {
		return nil
	}
}
