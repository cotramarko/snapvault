package engine

import (
	"fmt"
	"strings"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func configFromURL(url string) (DBConfig, error) {
	// `url` is a connection string to a postgres DB
	// of the form "postgres://user:pass@host:port/dbname"
	urlParts := strings.Split(url, "://")
	if len(urlParts) != 2 {
		return DBConfig{}, fmt.Errorf("invalid URL: %s", url)
	}

	urlParts = strings.Split(urlParts[1], "@")
	if len(urlParts) != 2 {
		return DBConfig{}, fmt.Errorf("invalid URL: %s", url)
	}

	credParts := strings.Split(urlParts[0], ":")
	if len(credParts) != 2 {
		return DBConfig{}, fmt.Errorf("invalid URL: %s", url)
	}

	hostParts := strings.Split(urlParts[1], ":")
	if len(hostParts) != 2 {
		return DBConfig{}, fmt.Errorf("invalid URL: %s", url)
	}

	portNamePart := strings.Split(hostParts[1], "/")
	if len(portNamePart) != 2 {
		return DBConfig{}, fmt.Errorf("invalid URL: %s", url)
	}

	dbConfig := DBConfig{
		User:     credParts[0],
		Password: credParts[1],
		Host:     hostParts[0],
		Port:     portNamePart[0],
		Name:     portNamePart[1],
	}

	return dbConfig, nil
}
