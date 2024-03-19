package commands

import (
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Restore(db engine.Engine, restoreName string) error {
	if err := db.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v\n", err)
		return err
	}

	fullDbName, err := db.GetSnap(restoreName)
	if err != nil {
		slog.Error("Failed to get snapshot: %v\n", err)
		return err
	}

	_, err = db.TerminateConnections()
	if err != nil {
		slog.Error("Failed to terminate connections: %v\n", err)
		return err
	}

	_, err = db.DisableTemplate()
	if err != nil {
		slog.Error("Failed to disable template: %v\n", err)
		return err
	}

	_, err = db.Drop(db.GetName())
	if err != nil {
		slog.Error("Failed to drop database: %v\n", err)
		return err
	}

	_, err = db.CreateFromSnap(fullDbName)
	if err != nil {
		slog.Error("Failed to create database from snapshot: %v\n", err)
		return err
	}

	return db.Close()
}
