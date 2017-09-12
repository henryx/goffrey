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
	Address string `names:"-a, --address" usage:"Address to get"`
	Rev     bool   `names:"-R, --rev" usage:"Reverse lookup" default:"false"`
}

func Get(log *logging.Logger, cfg *ini.File, data GetData) (string, error) {
	var db *sql.DB
	var hostexists bool
	var result string
	var err error

	if data.Address == "" {
		log.Debug("Section name is empty")
		return "", errors.New("No section name passed")
	}

	db, err = openDb(cfg)
	if err != nil {
		return "", err
	}
	defer db.Close()

	if !data.Rev {
		hostexists = checkHost(db, data.Section, data.Address)
		if !hostexists {
			return "", errors.New("Address " + data.Address + " not exists")
		}

		result, err = dbstore.GetIP(db, data.Section, data.Address)
		if err != nil {
			return "", err
		}
	} else {
		result, err = dbstore.GetHost(db, data.Section, data.Address)
		if err != nil {
			return "", err
		}

		if result == "" {
			return "", errors.New("Address has not associated hostname")
		}
	}

	return result, nil
}
