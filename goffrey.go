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
	"github.com/go-ini/ini"
	"github.com/cosiner/flag"
	"os"
	"os/user"
	"path/filepath"
	"actions"
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
	Cfg      string `names:"-c, --cfg" usage:"Set configuration file"`
	Register actions.RegisterData `usage:"Register a network"`
	Unregister struct {
		Enable bool
		Name   string `names:"-n, --name" usage:"Name of the network to unregister"`
	} `usage:"Unregister a network"`
}

func setCfg(cfg string) *ini.File {
	var filename string
	var res *ini.File
	var err error

	uid, _ := user.Current()

	if cfg != "" {
		filename = cfg
	} else {
		if _, err := os.Stat(uid.HomeDir + string(filepath.Separator) + ".goffreyrc"); os.IsNotExist(err) {
			filename = string(filepath.Separator) + "etc" + string(filepath.Separator) + "goffrey.cfg"
		} else {
			filename = uid.HomeDir + string(filepath.Separator) + ".goffreyrc"
		}
	}

	res, err = ini.Load([]byte{}, filename)
	if err != nil {
		fmt.Println("Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func main() {
	var args Args
	var cfg *ini.File
	testcode() // TODO: to remove

	set := flag.NewFlagSet(flag.Flag{})
	set.StructFlags(&args)
	set.Parse()

	if len(os.Args) == 1 {
		fmt.Println("No arguments passed")
		set.Help(false)
		os.Exit(0)
	}

	cfg = setCfg(args.Cfg)

	if args.Register.Enable {
		data := args.Register
		actions.Register(cfg, data)
	} else if args.Unregister.Enable {
		// TODO: implement this
	}
}
