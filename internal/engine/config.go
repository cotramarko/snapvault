package engine

import (
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

const ENV_NAME = "DATABASE_URL"
const TOML_FILE_NAME = "snapvault.toml"

func configFromEnv() (config DBConfig, exists bool) {
	val, exists := os.LookupEnv(ENV_NAME)
	if exists {
		config = configFromURL(val)
	}
	return
}

func configFromTOML(dir string) (config DBConfig, exists bool) {
	pathToFile := path.Join(dir, TOML_FILE_NAME)

	c := struct {
		URL string `toml:"url"`
	}{}
	_, err := toml.DecodeFile(pathToFile, &c)
	if err != nil {
		return
	}
	config = configFromURL(c.URL)
	exists = true
	return
}

func configFromURL(url string) DBConfig {
	// `url` is a connection string to a postgres DB
	// of the form "postgres://user:pass@host:port/dbname"
	urlParts := strings.Split(url, "://")
	if len(urlParts) != 2 {
		panic("Invalid connection string")
	}

	urlParts = strings.Split(urlParts[1], "@")
	if len(urlParts) != 2 {
		panic("Invalid connection string")
	}

	credParts := strings.Split(urlParts[0], ":")
	if len(credParts) != 2 {
		panic("Invalid connection string")
	}

	hostParts := strings.Split(urlParts[1], ":")
	if len(hostParts) != 2 {
		panic("Invalid connection string")
	}

	portNamePart := strings.Split(hostParts[1], "/")
	if len(portNamePart) != 2 {
		panic("Invalid connection string")
	}

	return DBConfig{
		User:     credParts[0],
		Password: credParts[1],
		Host:     hostParts[0],
		Port:     portNamePart[0],
		Name:     portNamePart[1],
	}
}
