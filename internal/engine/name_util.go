package engine

import "strings"

type DBname string
type SnapName string

const DB_NAME_SUFFIX = "_snapvault"

func ToDBname(snapName SnapName) DBname {
	return DBname(snapName + DB_NAME_SUFFIX)
}

func ToSnapName(dbName DBname) SnapName {
	return SnapName(strings.TrimSuffix(string(dbName), DB_NAME_SUFFIX))
}
