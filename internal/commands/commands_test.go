package commands

import (
	"fmt"
	"testing"

	"github.com/cotramarko/snapvault/internal/engine"
)

func setupTest() (*engine.Engine, func()) {
	conf := engine.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "acmeuser",
		Password: "acmepassword",
		Name:     "acmedb",
	}
	db := engine.NewEngine(conf)
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	return db, func() {
		db.Close()
	}
}

func TestSave(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	err := Save(*db, "edited3")
	if err != nil {
		t.Error(err)
	}
}

func TestRestore(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	err := Restore(*db, "fresh")
	if err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	res, err := List(*db)
	for _, r := range res {
		t.Logf("Row: %v", fmt.Sprintf("%v", r))
	}
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	err := Delete(*db, "edited3")
	if err != nil {
		t.Error(err)
	}
}
