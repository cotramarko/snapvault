package commands

import (
	"fmt"
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Restore(e engine.Engine, snapName engine.SnapName) error {
	if err := e.Connect(); err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to database: %v\n", err))
		return err
	}

	_, err := e.GetSnap(snapName)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get snapshot: %v\n", err))
		return err
	}

	_, err = e.TerminateConnections()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to terminate connections: %v\n", err))
		return err
	}

	_, err = e.DisableTemplate()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to disable template: %v\n", err))
		return err
	}

	_, err = e.Drop(engine.DBname(e.GetName())) // FIXME: should be cleaner
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to drop database: %v\n", err))
		return err
	}

	_, err = e.CreateFromSnap(snapName)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create database from snapshot: %v\n", err))
		return err
	}

	return e.Close()
}
