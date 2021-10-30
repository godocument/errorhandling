package dao

import (
	"context"
	"database/sql"
	"fmt"
)

type DAO interface {
	Get(ctxt context.Context, id int64) (*Project, error)
}

func New() DAO {
	return &dao{
		db: nil, // a placeholder only for this demo codelet
	}
}

type dao struct {
	db *sql.DB
}

func (d *dao) Get(ctxt context.Context, id int64) (*Project, error) {
	var project Project

	if err := d.db.QueryRow("SELECT id, name FROM project WHERE id = ?",
		id).Scan(&project.Id, &project.Name); err != nil {
		// do not return sql.ErrNoRows error
		// so I can hide the sql implementation
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project %d not exist", id)
		} else {
			return nil, fmt.Errorf("fail to get project: %v", err)
		}
	}

	return &project, nil
}
