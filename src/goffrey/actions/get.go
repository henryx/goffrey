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

type GetData struct {
	Enable  bool
	Section string `names:"-s, --section" usage:"Define the section to get"`
	Name    string `names:"-n, --name" usage:"Name of the host to get"`
}

func Get(log *logging.Logger, cfg *ini.File, data GetData) (string, error) {
	var db *sql.DB
	var hostexists bool
	var result string
	var err error

	if data.Name == "" {
		log.Debug("Section name is empty")
		return "", errors.New("No section name passed")
	}

	db, err = openDb(cfg)
	if err != nil {
		return "", err
	}
	defer db.Close()

	hostexists = checkHost(db, data.Section, data.Name)
	if !hostexists {
		return "", errors.New("Hostname " + data.Name + " not exists")
	}

	result, err = dbstore.GetIP(db, data.Section, data.Name)
	if err != nil {
		return "", err
	}

	return result, nil
}
