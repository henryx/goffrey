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

const help = `Usage: goffrey [-h|--help] <command>
	Global options are:
	  -h, --help: Print this help
	Commands are:
	  help: Print help of a command
	  add: Add network
	  del: Delete network
`

type Args struct {
	Action  string
	Name    string
	Network string
	Netmask int
}



func Parse() (Args, error) {
	res := Args{}
	if len(os.Args) <= 1 || utils.ContainStr(os.Args, "-h") || utils.ContainStr(os.Args, "--help") {
		fmt.Println(help)
		os.Exit(0)
	} else {
		if len(os.Args[2:]) <= 0 {
			return Args{}, errors.New("Not enough parameters (see \"-h\" option)")
		}

		switch os.Args[1] {
		case "help":
			parsehelpargs(os.Args[2])
		case "add":
			res.Action = "add"
			if err := parseaddargs(&res, os.Args[2:]); err != nil {
				return Args{}, err
			}
		case "del":
			res.Action = "del"
			parsedelargs(&res, os.Args[2:])
		}
	}

	return res, nil
}
