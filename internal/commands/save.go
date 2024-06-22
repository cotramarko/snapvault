package commands

import (
	"fmt"
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Save(e engine.Engine, snapName engine.SnapName) error {
	if err := e.Connect(); err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to database: %v\n", err))
		return err
	}

	_, err := e.TerminateConnections()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to terminate connections: %v\n", err))
		return err
	}

	_, err = e.EnableTemplate()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to enable template: %v\n", err))
		return err
	}

	_, err = e.Snap(snapName)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create snapshot: %v\n", err))
		return err
	}

	return e.Close()
}
