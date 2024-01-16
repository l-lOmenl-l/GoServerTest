package service

import (
	"example/web-servise-gin/internal/domain"
	"example/web-servise-gin/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
