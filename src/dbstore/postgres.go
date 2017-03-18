/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func OpenPostgres(host string, port int, user string, password string, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, password, host, port, database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
