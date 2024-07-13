package engine

import (
	"database/sql"
	"fmt"
	"github.com/cotramarko/snapvault/internal/connection"
)

type Engine struct {
	config DBConfig
	db     *sql.DB
}

func newEngine(config DBConfig) *Engine {
	return &Engine{config: config}
}

func DirectEngine(url string) (*Engine, error) {
	config, err := configFromURL(url)
	if err != nil {
		return nil, err
	}
	return newEngine(config), nil
}

func LoadEngine(dir string) (*Engine, error) {
	var (
		url    connection.URL
		config DBConfig
		err    error
	)

	url, err = connection.ProvideURL(dir)
	if err != nil {
		return nil, err
	}
	config, err = configFromURL(url.URL)
	if err != nil {
		return nil, fmt.Errorf("(%s) %v", url.Provider, err)
	}
	return newEngine(config), nil

}
