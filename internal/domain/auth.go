package domain

type SignIn struct {
	//Login    string `json:"login" binding:"required"`
	//Password string `json:"password" binding:"required"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	FIO         string `json:"fio"`
}
