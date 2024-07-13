package engine

import "strings"

type DBname string
type SnapName string

const DbNameSuffix = "_snapvault"

func ToDBname(snapName SnapName) DBname {
	return DBname(snapName + DbNameSuffix)
}

func ToSnapName(dbName DBname) SnapName {
	return SnapName(strings.TrimSuffix(string(dbName), DbNameSuffix))
}
