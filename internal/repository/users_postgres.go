package repository

import (
	"example/web-servise-gin/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) GetDetail(user_id int) (domain.UserDetail, error) {
	var userdetail domain.UserDetail

	query := fmt.Sprintf("SELECT lastname, firstname, fathername, email, users.active, superuser, staff, registrationsdate, entrydate, salon.name as \"salon\", cities.name as \"cities\" " +
		"FROM users " +
		"inner join salon on users.salon_id=salon.id " +
		"inner join cities on salon.city_id=cities.id " +
		"where users.id=$1")

	err := r.db.Get(&userdetail, query, user_id)
	return userdetail, err
}

func (r *UsersPostgres) GetAll() ([]domain.AllUsers, error) {

	var allusers []domain.AllUsers

	query := fmt.Sprintf("SELECT countries.id AS \"country_id\", countries.name AS \"country\", district.id AS \"district_id\"," +
		"district.name AS \"district\", regions.id AS \"region_id\", regions.name AS \"region\"," +
		"cities.id AS \"city_id\", cities.name AS \"city\", salon.id AS \"salon_id\", salon.name AS \"salon\", users.id AS \"user_id\", users.login AS \"user\" " +
		"FROM users " +
		"INNER JOIN salon ON users.salon_id=salon.id " +
		"INNER JOIN cities ON salon.id=cities.id " +
		"INNER JOIN regions ON regions.id=cities.id " +
		"INNER JOIN district ON district.id=regions.id " +
		"INNER JOIN countries ON countries.id=district.id")

	err := r.db.Select(&allusers, query)

	return allusers, err
}
