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
	"os"
	"flag"
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

func addargs(args []string) {
	action = "add"
	addcommand := flag.NewFlagSet("add", flag.ExitOnError)
	addcommand.StringVar(&section, "S", "", "Specify the section")

	addcommand.Parse(args)
}

func delargs(args []string) {
	action = "del"

	delcommand := flag.NewFlagSet("del", flag.ExitOnError)
	delcommand.Parse(args)
}

func contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func initargs() {
	var help = `Usage: goffrey [-h|--help] <command> [options]
	Global options are:
	  -h, --help: Print this help
	Commands are:
	  add: Add network
	  del: Delete network
	`
	if len(os.Args) <= 1 || contains(os.Args, "-h") || contains(os.Args, "--help") {
		fmt.Println(help)
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "add":
			addargs(os.Args[2:])
		case "del":
			delargs(os.Args[2:])
		}
	}
}

func main() {
	testcode()
	initargs()
}
