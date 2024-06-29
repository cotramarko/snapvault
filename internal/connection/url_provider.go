package connection

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

const EnvName = "DATABASE_URL"
const TomlFilename = "snapvault.toml"

type Provider string

const (
	tomlProvider Provider = TomlFilename
	envProvider  Provider = EnvName
)

type URL struct {
	URL      string
	Provider Provider
}

func urlFromTOML(dir string) (string, error) {
	pathToFile := path.Join(dir, TomlFilename)

	c := struct {
		URL string `toml:"url"`
	}{}
	_, err := toml.DecodeFile(pathToFile, &c)
	if err != nil {
		return "", err
	}
	return c.URL, nil
}

func tomlFileExists(dir string) bool {
	pathToFile := path.Join(dir, TomlFilename)
	if _, err := os.Stat(pathToFile); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(fmt.Sprintf("Error checking if %s exists: %v", pathToFile, err))
		}
	}
	return true
}

func ProvideURL(dir string) (URL, error) {
	if tomlFileExists(dir) {
		url, err := urlFromTOML(dir)
		if err != nil {
			return URL{}, fmt.Errorf("(%s) %v", tomlProvider, err)
		}
		return URL{URL: url, Provider: tomlProvider}, nil
	}
	url, exists := os.LookupEnv(EnvName)
	if exists {
		return URL{URL: url, Provider: envProvider}, nil
	}
	return URL{}, errors.New("no database URL provided")
}
