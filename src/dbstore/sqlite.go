/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func OpenSQLite(location string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}

	return db, nil
}