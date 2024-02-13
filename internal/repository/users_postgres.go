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
	query := fmt.Sprintf("SELECT lastname, firstname, fathername, email, users.active, superuser, staff, registrationsdate, entrydate, salon.name as \"salon_name\", cities.name as \"cities_name\" " +
		"FROM users " +
		"inner join salon on users.salon_id=salon.id " +
		"inner join cities on salon.city_id=cities.id " +
		"where users.id=$1")

	err := r.db.Get(&userdetail, query, user_id)

	return userdetail, err
}

func (r *UsersPostgres) GetAll() ([]domain.AllUsers, error) {

	var allusers []domain.AllUsers

	query := fmt.Sprintf("SELECT countries.name as \"country\", district.name as \"district\", regions.name as \"region\", cities.name as \"city\", salon.name as \"salon\", users.login as \"user\" " +
		"FROM countries " +
		"inner join district on district.country_id=countries.id " +
		"inner join regions on regions.district_id=district.id " +
		"inner join cities on cities.region_id=regions.id " +
		"inner join salon on salon.city_id = cities.id " +
		"inner join users on users.salon_id = salon.id ")

	err := r.db.Select(&allusers, query)

	return allusers, err
}
