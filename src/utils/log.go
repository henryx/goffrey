/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package utils

import (
	"io"
	"log"
)

type Log struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
}

func (l *Log) Init(infoHandle, warningHandle, errorHandle, debugHandle io.Writer) {
	// Usage:
	// l.Init(os.Stdout, os.Stdout, os.Stderr, ioutil.Discard)
	// l.Debug.Println("log message")

	l.Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	l.Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	l.Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)

	l.Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
