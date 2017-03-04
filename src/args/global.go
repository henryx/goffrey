/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package args

import (
	"os"
	"fmt"
	"errors"
	"utils"
)

const help = `Usage: goffrey [OPTIONS]... <COMMANDS>
	Global options are:
	  -c, --cfg CONF
	      Use CONF as configuration file. Defaults are /etc/goffrey.cfg or ~/.goffreyrc
	  -h, --help:
	      Print this help and exit
	Commands are:
	  help:
	      Print the help of a command
	  add:
	      Add a network
	  del:
	      Delete a network
`

type Args struct {
	Config  string
	Action  string
	Name    string
	Network string
	Netmask int
}

func Parse() (Args, error) {
	res := Args{}
	command := 1

	if len(os.Args) <= 1 || utils.ContainStr(os.Args, "-h") || utils.ContainStr(os.Args, "--help") {
		fmt.Println(help)
		os.Exit(0)
	} else {
		if len(os.Args[2:]) <= 0 {
			return Args{}, errors.New("Not enough parameters (see \"-h\" option)")
		}

		if utils.ContainStr(os.Args, "-c") ||utils.ContainStr(os.Args, "--cfg") {
			idx := -1
			if idx = utils.IndexStr(os.Args, "-c"); idx != -1 {
				command = idx + 1
			} else {
				idx = utils.IndexStr(os.Args, "--cfg")
			}
			res.Config = os.Args[idx + 1]
		}

		switch os.Args[command] {
		case "help":
			parsehelpargs(os.Args[command])
		case "add":
			res.Action = "add"
			if err := parseaddargs(&res, os.Args[command+1:]); err != nil {
				return Args{}, err
			}
		case "del":
			res.Action = "del"
			parsedelargs(&res, os.Args[command+1:])
		}
	}

	return res, nil
}
