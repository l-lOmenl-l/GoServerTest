package service

import (
	"example/web-servise-gin/internal/domain"
	"example/web-servise-gin/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetDetail(user_id int) (domain.UserDetail, error) {
	return s.repo.GetDetail(user_id)
}

func (s *UsersService) GetAll() ([]domain.Country, error) {
	return s.repo.GetAll()
}
