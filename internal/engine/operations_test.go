package engine

import (
	"fmt"
	"testing"
)

func setupTest() (*Engine, func()) {
	conf := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "acmeuser",
		Password: "acmepassword",
		Name:     "postgres", // connect to default postgres DB
	}
	db := NewEngine(conf)
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	return db, func() {
		db.db.Close()
	}
}

func TestDatabasePing(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	db.db.Ping()

	rows, err := db.db.Query(`
		SELECT
			datname
		FROM pg_database
		`,
	)
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r string
		rows.Scan(&r)
		t.Logf("Row: %v", r)
	}

	db.db.Close()
}

func TestDatabaseClone(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	// Kill all connections
	res, err := db.db.Exec(
		`SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'acmedb';`,
	)
	if err != nil {
		t.Error(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		t.Error(err)
	}
	if rowsAffected == 0 {
		t.Logf("No rows affected")
	}

	res, err = db.db.Exec(
		`ALTER DATABASE acmedb WITH is_template TRUE;`,
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.RowsAffected())

	res, err = db.db.Exec(
		`CREATE DATABASE acme_clone TEMPLATE acmedb;`,
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.RowsAffected())

	db.db.Close()

}

func TestDatabaseRestore(t *testing.T) {
	db, teardown := setupTest()
	defer teardown()

	// Kill all connections
	res, err := db.db.Exec(
		`SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'acmedb';`,
	)
	if err != nil {
		t.Error(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		t.Error(err)
	}
	if rowsAffected == 0 {
		t.Logf("No rows affected")
	}

	// Remove template tag from acmedb
	_, err = db.db.Exec(
		`ALTER DATABASE acmedb WITH is_template FALSE;`,
	)
	if err != nil {
		t.Error(err)
	}

	// Drop the database
	_, err = db.db.Exec(
		`DROP DATABASE acmedb;`,
	)
	if err != nil {
		t.Error(err)
	}

	// Create a new database from the template
	_, err = db.db.Exec(
		`CREATE DATABASE acmedb TEMPLATE acme_clone;`,
	)
	if err != nil {
		t.Error(err)
	}
}
