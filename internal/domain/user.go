package domain

type User struct {
	Id                int    `json:"-" db:"id"`
	Login             string `json:"login" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Lastname          string `json:"lastname" binding:"required"`
	Firstname         string `json:"firstname" binding:"required"`
	Fathername        string `json:"fathername" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Active            bool   `json:"active" binding:"required"`
	Superuser         bool   `json:"superuser" binding:"required"`
	Staff             bool   `json:"staff" binding:"required"`
	RegistrationsDate string `json:"registrationsDate" binding:"required"`
	EntryDate         string `json:"entryDate" binding:"required"`
	Salon             int    `json:"salon" binding:"required"`
}
