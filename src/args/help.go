/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package args

import "fmt"

const addcmd = `add <name> <network> <netmask>
    Add a network. Parameters are:
        <name>    - Specify a name for the network
        <network> - Specify the network addresses
        <netmask> - Specify the netmask
`

func parsehelpargs(arg string) {
	switch arg {
	case "add":
		fmt.Println(addcmd)
	default:
		fmt.Println("Unknown command: " + arg)
	}
}
