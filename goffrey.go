/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"actions"
	"github.com/cosiner/flag"
	"github.com/go-ini/ini"
	"github.com/op/go-logging"
	"ip"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func testcode(log *logging.Logger) {
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
		log.Error("Errorr: " + err.Error())
	}

	log.Debug(ips)
}

type Args struct {
	Cfg        string                 `names:"-c, --cfg" usage:"Set configuration file"`
	Quiet      bool                   `names:"-q, --quiet" usage:"Quiet mode"`
	Verbose    bool                   `names:"-v, --verbose" usage:"Verbose mode"`
	Register   actions.RegisterData   `usage:"Register a network"`
	Unregister actions.UnregisterData `usage:"Unregister a network"`
}

func setCfg(log *logging.Logger, cfg string) *ini.File {
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
		log.Error("Error about reading config file:", err)
		os.Exit(1)
	}

	return res
}

func setLog(level logging.Level, filename string) *logging.Logger {
	var backend *logging.LogBackend
	var log = logging.MustGetLogger("Goffrey")
	var format logging.Formatter

	if strings.ToUpper(level.String()) != "DEBUG" {
		format = logging.MustStringFormatter(
			"%{time:2006-01-02 15:04:05.000} %{level} - Goffrey - %{message}",
		)
	} else {
		format = logging.MustStringFormatter(
			"%{time:2006-01-02 15:04:05.000} %{level} - %{shortfile} - Goffrey - %{message}",
		)
	}

	if filename == "" {
		backend = logging.NewLogBackend(os.Stderr, "", 0)
	} else {
		fo, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		backend = logging.NewLogBackend(fo, "", 0)
	}

	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(level, "")

	logging.SetBackend(backendLeveled)
	logging.SetFormatter(format)

	return log
}

func register(log *logging.Logger, cfg *ini.File, data actions.RegisterData) {
	err := actions.Register(log, cfg, data)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Section %s registered", data.Name)
	}
}

func main() {
	var args Args
	var cfg *ini.File
	var level logging.Level

	set := flag.NewFlagSet(flag.Flag{})
	set.StructFlags(&args)
	set.Parse()

	if args.Verbose {
		level = logging.DEBUG
	} else if args.Quiet {
		level = logging.CRITICAL
	} else {
		level = logging.INFO
	}
	log := setLog(level, "")

	testcode(log) // TODO: to remove

	cfg = setCfg(log, args.Cfg)

	if args.Register.Enable {
		register(log, cfg, args.Register)
	} else if args.Unregister.Enable {
		// TODO: implement this
	} else {
		log.Error("No action passed")
		set.Help(false)
		os.Exit(0)
	}
}
