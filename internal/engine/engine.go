package engine

import (
	"database/sql"
)

type Engine struct {
	config DBConfig
	db     *sql.DB
}

func (e *Engine) GetName() string {
	return e.config.Name
}

func new(config DBConfig) *Engine {
	return &Engine{config: config}
}

func DirectEngine(url string) *Engine {
	return new(configFromURL(url))
}

func LoadEngine(dir string) *Engine {
	var (
		config DBConfig
		exists bool
	)
	config, exists = configFromTOML(dir)
	if exists {
		return new(config)
	}
	config, exists = configFromEnv()
	if exists {
		return new(config)
	}
	panic("No database configuration found")
}
