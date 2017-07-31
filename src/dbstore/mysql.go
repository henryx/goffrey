/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package dbstore

import "database/sql"

func OpenMySQL(host string, port int, user string, password string, database string) (*sql.DB, error) {
	// TODO: implement MySQL database connection
	return nil, nil
}
