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
	"utils"
)

type RegisterData struct {
	Enable  bool
	Name    string `names:"-n, --name" usage:"Set the name of the network"`
	Network string `names:"-N, --network" usage:"Set the network addresses"`
	Netmask string `names:"-M, --netmask" usage:"Set the network mask"`
}

func Register(log utils.Log, cfg *ini.File, data RegisterData) error {
	var db *sql.DB
	var err error

	if data.Name == "" {
		log.Println(utils.DEBUG, "Section name is empty")
		return errors.New("No section name passed")
	}

	dbtype := cfg.Section("general").Key("database").String()
	switch dbtype {
	case "sqlite":
		db, err = openSqlite(cfg.Section("sqlite").Key("location").String())
	case "postgres":
		sect := cfg.Section("postgres")
		db, err = openPg(sect)
	default:
		log.Println(utils.DEBUG, "Database specified: "+dbtype)
		return errors.New("Database not supported")
	}

	if err != nil {
		return err
	}
	defer db.Close()

	exists, err := dbstore.IsSectionExists(db, data.Name)
	if exists {
		return errors.New("Section " + data.Name + " already exists")
	} else if err != nil {
		log.Println(utils.DEBUG, "Error in check section: "+err.Error())
		return errors.New("Error about checking section")
	}

	err = dbstore.InsertSection(db, data.Name, data.Network, data.Netmask)
	if err != nil {
		log.Println(utils.DEBUG, "Error when insert section: "+err.Error())
		return errors.New("Error about insert section")
	} else {
		return nil
	}
}
