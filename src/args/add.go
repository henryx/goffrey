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

	if ok := net.ParseIP(args[0]); ok != nil {
		res.Network = args[1]
	} else {
		return errors.New("IP address not valid: " + args[0])
	}

	cidr := ip.ToCidr(args[2])
	res.Netmask = cidr

	return nil
}
