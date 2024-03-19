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

func (d *Engine) Drop() (int64, error) {
	execString := fmt.Sprintf("DROP DATABASE %s;", d.config.Name)
	res, err := d.db.Exec(execString)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
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
