package engine

import (
	"fmt"

	_ "github.com/lib/pq"
)

// TODO: Move these into separate structure
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
	execString := fmt.Sprintf("ALTER DATABASE %s WITH is_template TRUE;", d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) DisableTemplate() (int64, error) {
	execString := fmt.Sprintf("ALTER DATABASE %s WITH is_template FALSE;", d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) Snap(snapName string) (int64, error) {
	execString := fmt.Sprintf("CREATE DATABASE %s TEMPLATE %s;", snapName, d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) CreateFromSnap(snapName string) (int64, error) {
	execString := fmt.Sprintf("CREATE DATABASE %s TEMPLATE %s;", d.config.Name, snapName)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) Drop(name string) (int64, error) {
	execString := fmt.Sprintf("DROP DATABASE %s;", name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *Engine) GetSnapshots() ([]string, error) {
	rows, err := d.db.Query(`
	SELECT 
		regexp_substr(datname, '^[^_]+(?=_)'),
		pg_size_pretty(pg_database_size(pg_database.datname)), 
		(pg_stat_file('base/'|| oid ||'/PG_VERSION')).modification AS "created"
	FROM pg_database
		WHERE datname LIKE '%\_%'
	ORDER BY "created" DESC
	`)
	if err != nil {
		return nil, err
	}

	var snaps []string
	for rows.Next() {
		var name string
		var size string
		var created string
		if err := rows.Scan(&name, &size, &created); err != nil {
			return nil, err
		}
		snaps = append(snaps, name)
	}
	return snaps, nil
}

func (d *Engine) GetSnap(snapName string) (string, error) {
	rows, err := d.db.Query(`
		SELECT
			datname
		FROM pg_database
		WHERE
			pg_database.datname LIKE $1
		LIMIT 1;`, snapName+"_%",
	)

	if err != nil {
		return "", err
	}

	var fullDbName string
	for rows.Next() {
		if err := rows.Scan(&fullDbName); err != nil {
			return "", err
		}
	}

	return fullDbName, nil
}
