package commands

import (
	"fmt"
	"log/slog"

	"github.com/cotramarko/snapvault/internal/engine"
)

func List(e engine.Engine) ([]engine.SnapInfo, error) {
	if err := e.Connect(); err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to database: %v\n", err))
		return nil, err
	}

	return e.GetSnapshots()
}
