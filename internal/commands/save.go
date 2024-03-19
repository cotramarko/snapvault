package commands

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Save(db engine.Engine, saveName string) error {
	nameWithTS := fmt.Sprintf("%s_%s", saveName, time.Now().Format("20060102150405"))

	if err := db.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v\n", err)
		return err
	}

	_, err := db.TerminateConnections()
	if err != nil {
		slog.Error("Failed to terminate connections: %v\n", err)
		return err
	}

	_, err = db.EnableTemplate()
	if err != nil {
		slog.Error("Failed to enable template: %v\n", err)
		return err
	}

	_, err = db.Snap(nameWithTS)
	if err != nil {
		slog.Error("Failed to create snapshot: %v\n", err)
		return err
	}

	return db.Close()
}
