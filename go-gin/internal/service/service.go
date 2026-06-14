package service

import (
	models "github.com/Vasyl-Trefilov/bestBack/internal"
	"github.com/Vasyl-Trefilov/bestBack/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(user *models.User) (*models.User, error) {
	userId, err := s.repo.CreateUser(user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	user.ID = userId
	return user, nil
}
