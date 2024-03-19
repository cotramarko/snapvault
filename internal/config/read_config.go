package config

import (
	"github.com/cotramarko/snapvault/internal/engine"
)

func GetDefaultConfig() engine.DBConfig {
	return engine.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "acmeuser",
		Password: "acmepassword",
		Name:     "acmedb",
	}
}

func GetDefaultEngine() *engine.Engine {
	return engine.NewEngine(GetDefaultConfig())
}
