package main

import "database/sql"

type PragmaMap map[string]string

func defaultPragmaMap() PragmaMap {
	return PragmaMap{
		"busy_timeout":       "10000",
		"journal_mode":       "WAL",
		"journal_size_limit": "200000000",
		"synchronous":        "NORMAL",
		"foreign_keys":       "ON",
		"temp_store":         "MEMORY",
		"cache_size":         "-16000",
	}
}

func pragmaMapToDbUrl(pragmas PragmaMap) string {
	url := "?"
	for k, v := range pragmas {
		url += "_pragma=" + k + "(" + v + ")&"
	}
	url = url[:len(url)-1]
	return url
}

func connectToDb(dbPath string) error {
	pragmaMap := defaultPragmaMap()
	pragmas := pragmaMapToDbUrl(pragmaMap)
	_, err := sql.Open("sqlite", dbPath+pragmas)
	if err != nil {
		return err
	}
	return nil
}
