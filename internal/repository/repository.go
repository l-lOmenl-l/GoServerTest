package repository

import (
	"example/web-servise-gin/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(login, password string) (domain.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
