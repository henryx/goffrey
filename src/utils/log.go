/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package utils

import (
	"errors"
	"io"
	"log"
)

type loglevel int

type Log struct {
	level   loglevel
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	debug   *log.Logger
}

const (
	DEBUG loglevel = iota
	INFO
	WARNING
	ERROR
)

func (l *Log) Init(level loglevel, infoHandle, warningHandle, errorHandle, debugHandle io.Writer) {
	// Usage:
	// l.Init(os.Stdout, os.Stdout, os.Stderr, ioutil.Discard)
	// l.Debug.Println("log message")

	l.level = level

	l.info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	l.warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	l.error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)

	l.debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Log) Println(level loglevel, message ...interface{}) {
	switch level {
	case INFO:
		if l.level >= level {
			l.info.Println(message)
		}
	case WARNING:
		if l.level >= level {
			l.warning.Println(message)
		}
	case ERROR:
		if l.level >= level {
			l.error.Println(message)
		}
	case DEBUG:
		if l.level >= level {
			l.debug.Println(message)
		}
	}
}

func (l *Log) GetLogger(level loglevel) (*log.Logger, error) {
	switch level {
	case INFO:
		return l.info, nil
	case WARNING:
		return l.warning, nil
	case ERROR:
		return l.error, nil
	case DEBUG:
		return l.debug, nil
	default:
		return nil, errors.New("No logger specified")
	}
}
