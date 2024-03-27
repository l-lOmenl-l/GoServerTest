package domain

type User struct {
	Id                int    `json:"-" db:"id"`
	Login             string `json:"login" binding:"required" db:"login"`
	Password          string `json:"password" binding:"required" db:"password"`
	Lastname          string `json:"lastname" binding:"required" db:"lastname"`
	Firstname         string `json:"firstname" binding:"required" db:"firstname"`
	Fathername        string `json:"fathername" binding:"required" db:"fathername"`
	Email             string `json:"email" binding:"required" db:"email"`
	Active            bool   `json:"active" binding:"required" db:"active"`
	Superuser         bool   `json:"superuser" binding:"required" db:"superuser"`
	Staff             bool   `json:"staff" binding:"required" db:"staff"`
	RegistrationsDate string `json:"registrationsDate" binding:"required" db:"registrationsdate"`
	EntryDate         string `json:"entryDate" binding:"required" db:"entrydate"`
	Salon             int    `json:"salon" binding:"required" db:"salon_id"`
}

type UserDetail struct {
	Lastname          string `json:"lastname" binding:"required" db:"lastname"`
	Firstname         string `json:"firstname" binding:"required" db:"firstname"`
	Fathername        string `json:"fathername" binding:"required" db:"fathername"`
	Email             string `json:"email" binding:"required" db:"email"`
	Active            bool   `json:"active" binding:"required" db:"active"`
	Superuser         bool   `json:"superuser" binding:"required" db:"superuser"`
	Staff             bool   `json:"staff" binding:"required" db:"staff"`
	RegistrationsDate string `json:"registrationsDate" binding:"required" db:"registrationsdate"`
	EntryDate         string `json:"entryDate" binding:"required" db:"entrydate"`
	Salon             string `json:"salon" binding:"required" db:"salon_name"`
	City              string `json:"city" binding:"required" db:"cities_name"`
}

type AllUsers struct {
	Country_id  int    `json:"country_id" db:"country_id"`
	Country     string `json:"country" db:"country"`
	District_id int    `json:"district_id" db:"district_id"`
	District    string `json:"district" db:"district"`
	Region_id   string `json:"region_id" db:"region_id"`
	Region      string `json:"region" db:"region"`
	City_id     string `json:"city_id"  db:"city_id"`
	City        string `json:"city"  db:"city"`
	Salon_id    string `json:"salon_id" db:"salon_id"`
	Salon       string `json:"salon" db:"salon"`
	User_id     string `json:"user_id" db:"user_id"`
	User        string `json:"user" db:"user"`
}

type Salon map[string][]string
type City map[string][]Salon
type Region map[string][]City
type District map[string][]Region
type Country map[string][]District

type Country_ struct {
	Id   int    `json:"country_id" db:"id"`
	Name string `json:"country_name" db:"name"`
}

type District_ struct {
	Id   int    `json:"district_id" db:"id"`
	Name string `json:"country_name" db:"name"`
}

type Region_ struct {
	Id   int    `json:"regions_id" db:"id"`
	Name string `json:"country_name" db:"name"`
}

type City_ struct {
	Id   int    `json:"city_id" db:"id"`
	Name string `json:"country_name" db:"name"`
}

type Salon_ struct {
	Id   int    `json:"salon_id" db:"id"`
	Name string `json:"salon_name" db:"name"`
}

type User_ struct {
	Id    int    `json:"user_id" db:"id"`
	Login string `json:"user_login" db:"login"`
}
