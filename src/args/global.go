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
)

const help = `Usage: goffrey [-h|--help] <command>
	Global options are:
	  -h, --help: Print this help
	Commands are:
	  help: Print help of a command
	  add: Add network
	  del: Delete network
	`

func contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func Init() {
	if len(os.Args) <= 1 || contains(os.Args, "-h") || contains(os.Args, "--help") {
		fmt.Println(help)
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "help":
			helpargs(os.Args[2])
		case "add":
			addargs(os.Args[2:])
		case "del":
			delargs(os.Args[2:])
		}
	}
}