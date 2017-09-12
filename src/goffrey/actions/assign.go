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

type AssignData struct {
	Enable  bool
	Section string `names:"-s, --section" usage:"Define the section to assign"`
	Name    string `names:"-n, --name" usage:"Name of the host to assign"`
}

func Assign(log *logging.Logger, cfg *ini.File, data AssignData) (string, error) {
	var db *sql.DB
	var err error
	var sectexists, hostexists bool
	var ip string

	if data.Section == "" {
		log.Debug("Section name is empty")
		return "", errors.New("No section name passed")
	}

	db, err = openDb(cfg)
	if err != nil {
		return "", err
	}
	defer db.Close()

	sectexists, err = dbstore.IsSectionExists(db, data.Section)
	if err != nil {
		return "", err
	} else if !sectexists {
		return "", errors.New("Section " + data.Section + " not exists")
	}

	hostexists = checkHost(db, data.Section, data.Name)
	if hostexists {
		return "", errors.New("Hostname " + data.Name + " already assigned")
	}

	ip, err = dbstore.RetrieveFreeIP(db, data.Section)
	if err != nil {
		return "", err
	}

	err = dbstore.AssignHost(db, data.Section, data.Name, ip)
	if err != nil {
		return "", err
	}

	return ip, nil
}
