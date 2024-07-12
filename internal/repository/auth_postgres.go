package repository

import (
	"database/sql"
	"errors"
	"example/web-servise-gin/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func checkLogin(r *AuthPostgres, user domain.SignUp) error {
	query := fmt.Sprintf("SELECT id, login FROM %s WHERE login=$1", usersTable)
	var login string
	var id int

	row := r.db.QueryRow(query, user.Login)
	switch err := row.Scan(&id, &login); {
	case errors.Is(err, sql.ErrNoRows):
		return nil
	case err == nil:
		return fmt.Errorf("current user is already %w", err)
	default:
		return fmt.Errorf("unknow error %w", err)
	}
}

func (r *AuthPostgres) CreateUser(user domain.SignUp) (int, error) {
	var id int
	if checkLogin(r, user) != nil {
		return -1, errors.New("login already exists")
	}

	query := fmt.Sprintf("INSERT INTO %s (login, password, lastname, firstname, fathername, email, active, superuser, staff, registrationsDate, salon_id) "+
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Lastname, user.Firstname, user.Fathername, user.Email, true, false, false, time.Now().Format(time.RFC3339), user.Salon)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) SaveToken(token string) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (token) values ($1) RETURNING id", TokenTable)
	row := r.db.QueryRow(query, token)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) IdentifyToken(token string) error {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE token=$1", TokenTable)
	item := r.db.QueryRow(query, token)
	if err := item.Scan(&id); err != nil {
		return err
	}
	return nil
}
