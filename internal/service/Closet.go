package service

import (
	"example/web-servise-gin/internal/repository"
)

type ClosetService struct {
	repo repository.Closet
}

func NewClosetService(repo repository.Closet) *ClosetService {
	return &ClosetService{repo: repo}
}

func (s *ClosetService) AddTypeProduct(typeProduct string) (int, error) {
	return s.repo.AddTypeProduct(typeProduct)
}
