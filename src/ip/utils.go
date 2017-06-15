/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package ip

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ToCidr(mask string) (int, error) {
	var cidr int
	var err error

	if strings.Contains(mask, ".") {
		netmask := net.IPMask(net.ParseIP(mask).To4()) // If you have the mask as a string
		//netmask := net.IPv4Mask(255,255,255,0) // If you have the mask as 4 integer values

		cidr, _ = netmask.Size()
		if cidr == 0 {
			err := errors.New("No valid netmask: " + mask)
			return 0, err
		}

	} else {
		cidr, err = strconv.Atoi(mask)
		if err != nil {
			return 0, err
		}
	}

	return cidr, nil

}

func Range(network string, cidr int) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(network + "/" + strconv.Itoa(cidr))
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}
