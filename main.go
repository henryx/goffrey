/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package main

import (
	"fmt"
	"goffrey/actions"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cosiner/flag"
	"github.com/go-ini/ini"
	"github.com/op/go-logging"
)

type Args struct {
	Cfg        string                 `names:"-c, --cfg" usage:"Set configuration file"`
	Quiet      bool                   `names:"-q, --quiet" usage:"Quiet mode"`
	Verbose    bool                   `names:"-v, --verbose" usage:"Verbose mode"`
	Register   actions.RegisterData   `usage:"Register a network"`
	Unregister actions.UnregisterData `usage:"Unregister a network"`
	Assign     actions.AssignData     `usage:"Associate address"`
	Release    actions.ReleaseData    `usage:"Release associated address"`
	Get        actions.GetData        `usage:"Get associated address"`
}

func setCfg(log *logging.Logger, cfg string) *ini.File {
	var filename, homefile, systemfile string
	var res *ini.File
	var err error

	uid, _ := user.Current()
	homefile = uid.HomeDir + string(filepath.Separator) + ".goffreyrc"

	if runtime.GOOS == "windows" {
		systemfile = os.Getenv("ProgramData") + string(filepath.Separator) + "goffrey" + string(filepath.Separator) + "goffrey.cfg"
	} else {
		systemfile = string(filepath.Separator) + "etc" + string(filepath.Separator) + "goffrey.cfg"
	}

	if cfg != "" {
		filename = cfg
	} else {
		if _, err := os.Stat(homefile); os.IsNotExist(err) {
			if _, err := os.Stat(systemfile); os.IsNotExist(err) {
				log.Error("Configuration file does not exist in " + systemfile + " or " + homefile)
				os.Exit(1)
			} else {
				filename = systemfile
			}
		} else {
			filename = homefile
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

func unregister(log *logging.Logger, cfg *ini.File, data actions.UnregisterData) {
	err := actions.Unregister(log, cfg, data)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Section %s unregistered", data.Name)
	}
}

func assign(log *logging.Logger, cfg *ini.File, data actions.AssignData) {
	addr, err := actions.Assign(log, cfg, data)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Address %s assigned for host %s", addr, data.Name)
	}

}

func release(log *logging.Logger, cfg *ini.File, data actions.ReleaseData) {
	addr, err := actions.Release(log, cfg, data)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Address %s released for host %s", addr, data.Name)
	}
}

func get(log *logging.Logger, cfg *ini.File, data actions.GetData) {
	addr, err := actions.Get(log, cfg, data)
	if err != nil {
		log.Error(err.Error())
	} else {
		fmt.Printf("%s - %s\n", data.Address, addr)
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

	cfg = setCfg(log, args.Cfg)

	if args.Register.Enable {
		register(log, cfg, args.Register)
	} else if args.Unregister.Enable {
		unregister(log, cfg, args.Unregister)
	} else if args.Assign.Enable {
		assign(log, cfg, args.Assign)
	} else if args.Release.Enable {
		release(log, cfg, args.Release)
	} else if args.Get.Enable {
		get(log, cfg, args.Get)
	} else {
		log.Error("No action passed")
		set.Help(false)
		os.Exit(0)
	}
}
