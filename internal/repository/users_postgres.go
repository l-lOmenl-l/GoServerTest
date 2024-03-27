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

func (r *UsersPostgres) GetAll() ([]domain.Country, error) {
	var all []domain.Country

	var allcountry []domain.Country_
	var alldistrict []domain.District_
	var allregions []domain.Region_
	var allcity []domain.City_
	var allsalon []domain.Salon_
	var alluser []domain.User_

	salon_ := domain.Salon{}
	city_ := domain.City{}
	region_ := domain.Region{}
	district_ := domain.District{}
	country_ := domain.Country{}

	query := fmt.Sprintf("SELECT id, name FROM countries")
	err := r.db.Select(&allcountry, query)
	if err != nil {
		return all, err
	}

	for _, country := range allcountry {
		query = fmt.Sprintf("SELECT id, name FROM district where country_id=$1")
		err = r.db.Select(&alldistrict, query, country.Id)
		if err != nil {
			return all, err
		}

		for _, district := range alldistrict {
			query = fmt.Sprintf("SELECT id, name FROM regions where district_id=$1")
			err = r.db.Select(&allregions, query, district.Id)
			if err != nil {
				return all, err
			}

			for _, region := range allregions {
				query = fmt.Sprintf("SELECT id, name FROM cities where region_id=$1")
				err = r.db.Select(&allcity, query, region.Id)
				if err != nil {
					return all, err
				}

				for _, city := range allcity {
					query = fmt.Sprintf("SELECT id, name FROM salon where city_id=$1")
					err = r.db.Select(&allsalon, query, city.Id)
					if err != nil {
						return all, err
					}

					for _, salon := range allsalon {
						query = fmt.Sprintf("SELECT id, login FROM users where salon_id=$1")
						err = r.db.Select(&alluser, query, salon.Id)
						if err != nil {
							return all, err
						}

						for _, user := range alluser {
							salon_[salon.Name] = append(salon_[salon.Name], user.Login)
						}

						city_[city.Name] = append(city_[city.Name], salon_)
						salon_ = domain.Salon{}
					}

				}

				region_[region.Name] = append(region_[region.Name], city_)
				city_ = domain.City{}

			}

			district_[district.Name] = append(district_[district.Name], region_)
			region_ = domain.Region{}
		}
		country_[country.Name] = append(country_[country.Name], district_)
		district_ = domain.District{}

		all = append(all, country_)
		country_ = domain.Country{}
	}

	return all, err
}
