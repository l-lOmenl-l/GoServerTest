package service

import (
	"example/web-servise-gin/internal/domain"
	"example/web-servise-gin/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Users interface {
	GetDetail(user_id int) (domain.UserDetail, error)
	GetAll() ([]domain.AllUsers, error)
}

type Closet interface {
	AddTypeProduct(typeProduct string) (int, error)
}

type Service struct {
	Authorization
	Users
	Closet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUsersService(repos.Users),
		Closet:        NewClosetService(repos.Closet),
	}
}
