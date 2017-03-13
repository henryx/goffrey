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
	"github.com/cosiner/flag"
)

func testcode() {
	// TODO: remove it
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

type Args struct {
	Cfg string `names:"-c, --cfg" usage:"Set configuration file"`
	Add struct {
		Enable  bool
		Name    string `usage:"Set the name of the network"`
		Network string `usage:"Set the network addresses"`
		Netmask string `usage:"Set the network mask"`
	} `usage:"Add a network"`
	Del struct {
		Enable bool
		Name   string `usage:"Set the name of the network"`
	} `usage:"Delete a network"`
}

func main() {
	var args Args
	testcode()

	set := flag.NewFlagSet(flag.Flag{})
	set.StructFlags(&args)
	set.Parse()
	fmt.Println(args.Add.Name)
	set.Help(false)
}
