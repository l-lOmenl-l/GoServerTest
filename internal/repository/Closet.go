package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ClosetPostgres struct {
	db *sqlx.DB
}

func NewClosetPostgres(db *sqlx.DB) *ClosetPostgres {
	return &ClosetPostgres{db: db}
}

func (r *ClosetPostgres) AddTypeProduct(typeProduct string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", typeProductTable)
	row := r.db.QueryRow(query, typeProduct)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
