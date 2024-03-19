package commands

import (
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Delete(db engine.Engine, snapName string) error {
	if err := db.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v\n", err)
		return err
	}

	fullDbName, err := db.GetSnap(snapName)
	if err != nil {
		slog.Error("Failed to get snapshot: %v\n", err)
		return err
	}

	_, err = db.Drop(fullDbName)
	if err != nil {
		slog.Error("Failed to drop database: %v\n", err)
		return err
	}

	return db.Close()
}
