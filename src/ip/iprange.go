/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package ip

import "net"

func ToCidr(mask string) int {
	netmask := net.IPMask(net.ParseIP(mask).To4()) // If you have the mask as a string
	//netmask := net.IPv4Mask(255,255,255,0) // If you have the mask as 4 integer values

	prefix, _ := netmask.Size()
	return prefix

}