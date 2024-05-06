package repository

import (
	"example/web-servise-gin/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(login, password string) (domain.User, error)
	SaveToken(token string) error
	IdentifyToken(token string) error
}

type Users interface {
	GetDetail(id int) (domain.UserDetail, error)
	GetAll() ([]domain.AllUsers, error)
}

type Closet interface {
	AddTypeProduct(typeProduct string) (int, error)
}

type Repository struct {
	Authorization
	Users
	Closet
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUsersPostgres(db),
		Closet:        NewClosetPostgres(db),
	}
}
