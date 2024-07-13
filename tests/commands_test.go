package commands_test

import (
	"testing"

	"github.com/cotramarko/snapvault/internal/commands"
	"github.com/cotramarko/snapvault/internal/engine"
)

// Integration test, requires DB via docker-compose to be running
func setupTest() *engine.Engine {
	url := "postgres://acmeuser:acmepassword@localhost:5432/acmedb"
	e, err := engine.DirectEngine(url)
	if err != nil {
		panic(err)
	}
	err = e.Connect()
	if err != nil {
		panic(err)
	}
	return e
}

func TestCoreCommands(t *testing.T) {
	e := setupTest()

	err := commands.Save(*e, "fresh")
	if err != nil {
		t.Error(err)
	}
	t.Log("Saved `fresh`")

	err = commands.Restore(*e, "fresh")
	if err != nil {
		t.Error(err)
	}
	t.Log("Restored `fresh`")

	res, err := commands.List(*e)
	if len(res) != 1 {
		t.Error("Expected 1 snapshot, got", len(res))
	}

	if err != nil {
		t.Error(err)
	}
	t.Log("Listed snapshots")

	err = commands.Delete(*e, "fresh")
	if err != nil {
		t.Error(err)
	}
	t.Log("Deleted `fresh`")
}
