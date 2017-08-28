/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package actions

type RemoveData struct {
	Enable bool
	Name   string `names:"-n, --name" usage:"Name of the host to remove"`
}
