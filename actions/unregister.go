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
	"github.com/op/go-logging"
	"goffrey/dbstore"
)

type UnregisterData struct {
	Enable bool
	Name   string `names:"-n, --name" usage:"Name of the network to unregister"`
}

func Unregister(log *logging.Logger, cfg *ini.File, data UnregisterData) error {
	var db *sql.DB
	var err error

	if data.Name == "" {
		log.Debug("Section name is empty")
		return errors.New("No section name passed")
	}

	db, err = openDb(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	exists, err := dbstore.IsSectionExists(db, data.Name)
	if !exists {
		return errors.New("Section " + data.Name + " does not exists")
	} else if err != nil {
		log.Debug("Error in check section: " + err.Error())
		return errors.New("Error about checking section")
	}

	if err = dbstore.RemoveSection(db, data.Name); err != nil {
		log.Debug("Error when remove section: " + err.Error())
		return errors.New("Error about remove section")
	}
	return nil
}
