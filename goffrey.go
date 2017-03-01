/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
	"ip"
	"args"
)

func testcode() {
	var cidr string

	network := "192.168.0.0/255.255.255.252"
	mask := strings.Split(network, "/")

	if strings.Contains(mask[1], ".") {
		cidr = strconv.Itoa(ip.ToCidr(mask[1]))
	} else {
		cidr = mask[1]
	}

	ips, err := ip.Range(mask[0] + "/" + cidr)
	if err != nil {
		fmt.Println("Errorr: " + err.Error())
	}

	fmt.Println(ips)
}

var (
	action  string
	section string
)

func main() {
	testcode()
	args.Init()
}
