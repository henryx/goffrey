/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package args

import (
	"net"
	"errors"
	"ip"
)

func parseaddargs(res *Args, args []string) error {
	res.Name = args[0]

	if ok := net.ParseIP(args[1]); ok != nil {
		res.Network = args[1]
	} else {
		return errors.New("Network is not valid: " + args[1])
	}

	if cidr := ip.ToCidr(args[2]); cidr != 0 {
		res.Netmask = cidr
	} else {
		return errors.New("Netmask is not valid: " + args[2])

	}

	return nil
}
