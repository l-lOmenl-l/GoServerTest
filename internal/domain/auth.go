package domain

type SignIn struct {
	Place       string `json:"place" binding:"required"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	FIO         string `json:"fio"`
}

type SignUp struct {
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Lastname   string `json:"lastname" binding:"required" db:"lastname"`
	Firstname  string `json:"firstname" binding:"required" db:"firstname"`
	Fathername string `json:"fathername" binding:"required" db:"fathername"`
	Email      string `json:"email" binding:"required" db:"email"`
	Salon      int    `json:"salon" binding:"required" db:"salon_id"`
}
