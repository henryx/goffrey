/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

import (
	"github.com/op/go-logging"
	"github.com/go-ini/ini"
)

type AssignData struct {
	Enable bool
	Name   string `names:"-n, --name" usage:"Name of the host to assign"`
}

func Assign(log *logging.Logger, cfg *ini.File, data AssignData) (string, error) {
	var result string

	// TODO: implement this

	return result, nil
}