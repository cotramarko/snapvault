package engine

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
type Engine struct {
	config DBConfig
	db     *sql.DB
}

func NewEngine(config DBConfig) *Engine {
	return &Engine{config: config}
}

func (d *Engine) GetName() string {
	return d.config.Name
}

func (d *Engine) Connect() error {
	// FIXME: Find better way to inject "postgres"
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.config.Host, d.config.Port, d.config.User, d.config.Password, "postgres",
	)

	// TODO: Add a ping to check if connection is still alive
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

func (d *Engine) Close() error {
	return d.db.Close()
}
