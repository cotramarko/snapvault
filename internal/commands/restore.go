package commands

import (
	"fmt"

	"github.com/cotramarko/snapvault/internal/engine"
)

func Restore(db engine.Engine, restoreName string) error {
	if err := db.Connect(); err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return err
	}

	fullDbName, err := db.GetSnap(restoreName)
	if err != nil {
		fmt.Printf("Failed to get snapshot: %v\n", err)
		return err
	}

	_, err = db.TerminateConnections()
	if err != nil {
		fmt.Printf("Failed to terminate connections: %v\n", err)
		return err
	}

	_, err = db.DisableTemplate()
	if err != nil {
		fmt.Printf("Failed to disable template: %v\n", err)
		return err
	}

	_, err = db.Drop()
	if err != nil {
		fmt.Printf("Failed to drop database: %v\n", err)
		return err
	}

	_, err = db.CreateFromSnap(fullDbName)
	if err != nil {
		fmt.Printf("Failed to create database from snapshot: %v\n", err)
		return err
	}

	return db.Close()
}
