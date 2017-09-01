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
)

type AssignData struct {
	Enable bool
	Name   string `names:"-n, --name" usage:"Name of the host to assign"`
}

func Assign(log *logging.Logger, cfg *ini.File, data AssignData) (string, error) {
	var db *sql.DB
	var err error
	var result string

	if data.Name == "" {
		log.Debug("Section name is empty")
		return "", errors.New("No section name passed")
	}

	db, err = openDb(cfg)
	if err != nil {
		return "", err
	}
	defer db.Close()

	// TODO: implement this

	return result, nil
}
