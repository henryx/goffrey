/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"actions"
	"encoding/json"
	"github.com/cosiner/flag"
	"github.com/go-ini/ini"
	"ip"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"utils"
)

func testcode(log utils.Log) {
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
		log.Println(utils.ERROR, "Errorr: "+err.Error())
	}

	log.Println(utils.DEBUG, ips)
}

type Args struct {
	Cfg        string               `names:"-c, --cfg" usage:"Set configuration file"`
	Quiet      bool                 `names:"-q, --quiet" usage:"Quiet mode"`
	Verbose    bool                 `names:"-v, --verbose" usage:"Verbose mode"`
	Register   actions.RegisterData `usage:"Register a network"`
	Unregister struct {
		Enable bool
		Name   string `names:"-n, --name" usage:"Name of the network to unregister"`
	} `usage:"Unregister a network"`
}

func setCfg(log utils.Log, cfg string) *ini.File {
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
		log.Println(utils.ERROR, "Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func register(log utils.Log, cfg *ini.File, data actions.RegisterData) {
	err := actions.Register(log, cfg, data)
	if err != nil {
		log.Println(utils.ERROR, err.Error())
	}
}

func main() {
	var args Args
	var cfg *ini.File

	set := flag.NewFlagSet(flag.Flag{})
	set.StructFlags(&args)
	set.Parse()

	log := utils.Log{}
	if args.Verbose {
		log.Init(utils.DEBUG, os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	} else if args.Quiet {
		log.Init(utils.CRITICAL, os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	} else {
		log.Init(utils.ERROR, os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	}

	testcode(log) // TODO: to remove

	cfg = setCfg(log, args.Cfg)

	if args.Register.Enable {
		register(log, cfg, args.Register)
	} else if args.Unregister.Enable {
		// TODO: implement this
	} else {
		log.Println(utils.ERROR, "No action passed")
		set.Help(false)
		os.Exit(0)
	}
}
