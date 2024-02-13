package repository

import (
	"example/web-servise-gin/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password, lastname, firstname, fathername, email, active, superuser, staff, registrationsDate, entryDate, salon_id) "+
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Lastname, user.Firstname, user.Fathername, user.Email, user.Active, user.Superuser, user.Staff, user.RegistrationsDate, user.EntryDate, user.Salon)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
