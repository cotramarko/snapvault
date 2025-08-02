package engine

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
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
	createRes, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}

	description, err := d.GetCommentForDatabase(d.GetName())
	if err != nil {
		return 0, err
	}
	if description != nil {
		err = d.WriteCommentForDatabase(string(dbName), *description)
		if err != nil {
			return 0, err
		}
	}
	return createRes.RowsAffected()
}

func (d *Engine) CreateFromSnap(snapName SnapName) (int64, error) {
	dbName := ToDBname(snapName)
	description, err := d.GetCommentForDatabase(string(dbName))
	if err != nil {
		return 0, err
	}

	execString := fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s";`, d.GetName(), dbName)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}

	if description != nil {
		err = d.WriteCommentForDatabase(d.config.Name, *description)
		if err != nil {
			return 0, err
		}
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
	Comment  string
}

func (d *Engine) GetSnapshots() ([]SnapInfo, error) {
	rows, err := d.db.Query(`
	SELECT
		datname AS "snap_name",
		pg_size_pretty(pg_database_size(pg_database.datname)) AS "size",
		(pg_stat_file('base/'|| oid ||'/PG_VERSION')).modification AS "created",
		COALESCE(pg_shdescription.description, '') AS "comment"
	FROM pg_database
		LEFT JOIN pg_shdescription ON pg_shdescription.objoid = pg_database.oid
	WHERE pg_database.datname LIKE $1
	ORDER BY "created" DESC
	`, "%"+DbNameSuffix)
	if err != nil {
		return nil, err
	}

	var snaps []SnapInfo
	for rows.Next() {
		var dbName, size, created, comment string
		if err = rows.Scan(&dbName, &size, &created, &comment); err != nil {
			return nil, err
		}
		snaps = append(snaps, SnapInfo{
			SnapName: ToSnapName(DBname(dbName)),
			Size:     size,
			Created:  created,
			Comment:  comment,
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

func (e *Engine) GetCommentForDatabase(dbName string) (*string, error) {
	var description *string
	err := e.db.QueryRow(`
	SELECT
		description
	FROM pg_shdescription
	JOIN pg_database ON objoid = pg_database.oid
	WHERE datname = $1`, dbName).Scan(&description)
	if err == sql.ErrNoRows {
		// No comment is set, this is fine
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return description, nil
}

func (e *Engine) WriteCommentForDatabase(dbName string, comment string) error {
	_, err := e.db.Exec(fmt.Sprintf(`COMMENT ON DATABASE "%s" IS %s`, dbName, pq.QuoteLiteral(comment)))
	if err != nil {
		return err
	}

	return nil
}
