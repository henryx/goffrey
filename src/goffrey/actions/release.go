/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

import (
	"github.com/go-ini/ini"
	"github.com/op/go-logging"
)

type ReleaseData struct {
	Enable  bool
	Section string `names:"-s, --section" usage:"Define the section to assign"`
	Name    string `names:"-n, --name" usage:"Name of the host to release"`
}

func Release(log *logging.Logger, cfg *ini.File, data ReleaseData) (string, error) {
	// TODO: implement function

	return "", nil
}
