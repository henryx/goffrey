/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Goffrey
   Description   A simple IPAM
   License       GPL version 2 (see GPL.txt for details)
*/

package utils

func IndexStr(slice []string, element string) (int) {
	for pos, a := range slice {
		if a == element {
			return pos
		}
	}

	return -1
}

func ContainStr(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
