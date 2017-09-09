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

type GetData struct {
	Enable  bool
	Section string `names:"-s, --section" usage:"Define the section to get"`
	Name    string `names:"-n, --name" usage:"Name of the host to get"`
}

func Get(log *logging.Logger, cfg *ini.File, data GetData) (string, error) {
	// TODO implement function

	return "", nil
}