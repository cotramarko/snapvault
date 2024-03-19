package commands

import (
	"fmt"
	"time"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Save(db engine.Engine, saveName string) error {
	nameWithTS := fmt.Sprintf("%s_%s", saveName, time.Now().Format("20060102150405"))

	if err := db.Connect(); err != nil {
		return err
	}

	_, err := db.TerminateConnections()
	if err != nil {
		return err
	}

	_, err = db.EnableTemplate()
	if err != nil {
		return err
	}

	_, err = db.Snap(nameWithTS)
	if err != nil {
		return err
	}

	return db.Close()
}
