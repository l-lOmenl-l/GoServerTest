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
	Country  string `json:"country" db:"country"`
	District string `json:"district" db:"district"`
	Region   string `json:"region" db:"region"`
	City     string `json:"city"  db:"city"`
	Salon    string `json:"salon" db:"salon"`
	User     string `json:"user" db:"user"`
}
