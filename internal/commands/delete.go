package commands

import (
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Delete(e engine.Engine, snapName engine.SnapName) error {
	if err := e.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v\n", err)
		return err
	}

	dbName, err := e.GetSnap(snapName)
	if err != nil {
		slog.Error("Failed to get snapshot: %v\n", err)
		return err
	}

	_, err = e.Drop(dbName)
	if err != nil {
		slog.Error("Failed to drop database: %v\n", err)
		return err
	}

	return e.Close()
}
