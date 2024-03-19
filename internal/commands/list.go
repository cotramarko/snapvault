package commands

import (
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func List(db engine.Engine) ([]string, error) {
	if err := db.Connect(); err != nil {
		slog.Error("Failed to connect to database: %v\n", err)
		return nil, err
	}

	return db.GetSnapshots()
}
