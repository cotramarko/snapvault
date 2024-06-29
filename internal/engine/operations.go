package engine

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func (e *Engine) Connect() error {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.config.Host, e.config.Port, e.config.User, e.config.Password, "postgres",
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	e.db = db

	return nil
}

func (e *Engine) Close() error {
	return e.db.Close()
}

func (d *Engine) TerminateConnections() (int64, error) {
	res, err := d.db.Exec(
		`SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = $1;`,
		d.config.Name,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) EnableTemplate() (int64, error) {
	execString := fmt.Sprintf(`ALTER DATABASE "%s" WITH is_template TRUE;`, d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) DisableTemplate() (int64, error) {
	execString := fmt.Sprintf(`ALTER DATABASE "%s" WITH is_template FALSE;`, d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) Snap(snapName SnapName) (int64, error) {
	dbName := ToDBname(snapName)
	execString := fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s";`, dbName, d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) CreateFromSnap(snapName SnapName) (int64, error) {
	dbName := ToDBname(snapName)
	execString := fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s";`, d.config.Name, dbName)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) Drop(dbName DBname) (int64, error) {
	execString := fmt.Sprintf(`DROP DATABASE "%s";`, dbName)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

type SnapInfo struct {
	SnapName SnapName
	Size     string
	Created  string
}

func (d *Engine) GetSnapshots() ([]SnapInfo, error) {
	rows, err := d.db.Query(`
	SELECT 
		datname,
		pg_size_pretty(pg_database_size(pg_database.datname)), 
		(pg_stat_file('base/'|| oid ||'/PG_VERSION')).modification AS "created"
	FROM pg_database
		WHERE datname LIKE $1
	ORDER BY "created" DESC
	`, "%"+DbNameSuffix)
	if err != nil {
		return nil, err
	}

	var snaps []SnapInfo
	for rows.Next() {
		var dbName string
		var size string
		var created string
		if err = rows.Scan(&dbName, &size, &created); err != nil {
			return nil, err
		}
		snaps = append(snaps, SnapInfo{
			SnapName: ToSnapName(DBname(dbName)),
			Size:     size,
			Created:  created,
		})
	}
	return snaps, nil
}

func (d *Engine) GetSnap(snapName SnapName) (DBname, error) {
	dbName := ToDBname(snapName)
	rows, err := d.db.Query(`
		SELECT
			datname
		FROM pg_database
		WHERE
			pg_database.datname LIKE $1
		LIMIT 1;`, dbName,
	)

	if err != nil {
		return "", err
	}

	var fullDbName string
	for rows.Next() {
		if err = rows.Scan(&fullDbName); err != nil {
			return "", err
		}
	}

	if fullDbName == "" {
		return "", errors.New(fmt.Sprintf("Could not find snapshot with name %v", snapName))
	}
	return DBname(fullDbName), nil
}

func (e *Engine) GetName() string {
	return e.config.Name
}
